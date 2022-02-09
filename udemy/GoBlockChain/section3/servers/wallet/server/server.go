package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"text/template"

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

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename) // Path to this file
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer( http.Dir(path.Join(dir, assetsDir)))))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.port)), nil))
}

func NewWallServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}
