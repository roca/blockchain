package server

import (
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

const tempDir = "section3/servers/wallet/server/templates/"

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
		t, e := template.ParseFiles(path.Join(tempDir, "index.html"))
		if e != nil {
			log.Panicf("ERROR: %v", e)
		}
		t.Execute(w, nil)
	default:
		log.Printf("ERROR: Invalid request method: %v", req.Method)
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.port)), nil))
}

func NewWallServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}
