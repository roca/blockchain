package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"text/template"

	"udemy.com/goblockchain/section3/blockchain"
	"udemy.com/goblockchain/section3/utils"
	"udemy.com/goblockchain/section3/wallet"
)

const tempDir = "templates/"
const assetsDir = "assets/"

type WalletServer struct {
	port    uint16
	gateway string
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Dir(filename) // Path to this file
		t, e := template.ParseFiles(path.Join(dir, tempDir, "index.html"))
		if e != nil {
			log.Panicf("ERROR: %v", e)
		}
		e = t.Execute(w, nil)
		if e != nil {
			log.Panicf("ERROR: %v", e)
		}
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := json.Marshal(myWallet)
		_, e := io.WriteString(w, string(m[:]))
		if e != nil {
			log.Printf("ERROR: %v", e)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var wtr wallet.TransactionRequest
		e := decoder.Decode(&wtr)
		if e != nil {
			log.Printf("ERROR: %v", e)
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request")))
			return
		}
		if !wtr.Validate() {
			log.Println("ERROR: missing fields")
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request: missing fields")))
			return
		}

		publicKey := utils.PublicKeyFromString(*wtr.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*wtr.SenderPrivateKey, publicKey)
		// log.Printf("VALUE: %v",*wtr.Value)
		value, err := strconv.ParseFloat(*wtr.Value, 32)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("failed")))
			return
		}
		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")

		wt := wallet.NewTransaction(
			privateKey,
			publicKey,
			*wtr.SenderBlockchainAddress,
			*wtr.RecipientBlockchainAddress,
			value32,
		)
		signature := wt.GenerateSignature()
		signatureStr := signature.String()

		btr := &blockchain.TransactionRequest{
			SenderBlockchainAddress:    wtr.SenderBlockchainAddress,
			RecipientBlockchainAddress: wtr.RecipientBlockchainAddress,
			SenderPublicKey:            wtr.SenderPublicKey,
			Value:                      &value32,
			Signature:                  &signatureStr,
		}

		m, _ := json.Marshal(btr)
		buf := bytes.NewBuffer(m)

		resp, _ := http.Post(ws.Gateway()+"/transactions", "application/json", buf)
		if resp.StatusCode == 201 {
			io.WriteString(w, string(utils.JsonStatus("success")))
			return
		}
		io.WriteString(w, string(utils.JsonStatus("failed")))

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
	}
}

func (ws *WalletServer) WalletAmount(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		blockchainAddress := req.URL.Query().Get("blockchain_address")
		endpoint := fmt.Sprintf("%s/amount", ws.Gateway())

		client := &http.Client{}
		bcsReq, _ := http.NewRequest("GET", endpoint, nil)
		q := bcsReq.URL.Query()
		q.Add("blockchain_address", blockchainAddress)
		bcsReq.URL.RawQuery = q.Encode()

		bcsResp, err := client.Do(bcsReq)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("HTTP GET wallet amount failed")))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		if bcsResp.StatusCode == 200 {
			decoder := json.NewDecoder(bcsResp.Body)
			var bar blockchain.AmountResponse
			err := decoder.Decode(&bar)
			if err != nil {
				log.Printf("ERROR: %v", err)
				io.WriteString(w, string(utils.JsonStatus("Decoding GET amount failed")))
				return
			}
			m, _ := json.Marshal(struct {
				Message string  `json:"message"`
				Amount  float32 `json:"amount"`
			}{
				Message: "success",
				Amount:  bar.Amount,
			})
			io.WriteString(w, string(m[:]))
		} else {
			io.WriteString(w, string(utils.JsonStatus("StatusCode GET amount not 200")))
		}
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/wallet/amount", ws.WalletAmount)
	http.HandleFunc("/transaction", ws.CreateTransaction)

	// Static assets
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename) // Path to this file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(path.Join(dir, assetsDir)))))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.port)), nil))
}

func NewWallServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}
