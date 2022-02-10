package server

import (
	"bytes"
	"encoding/json"
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
		log.Printf("ERROR: Invalid request method: %v", req.Method)
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
		log.Printf("ERROR: Invalid request method: %v", req.Method)
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
		log.Printf("ERROR: Invalid request method: %v", req.Method)
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
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
