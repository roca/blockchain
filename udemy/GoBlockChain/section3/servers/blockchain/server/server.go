package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"udemy.com/goblockchain/section3/blockchain"
	"udemy.com/goblockchain/section3/utils"
	"udemy.com/goblockchain/section3/wallet"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainServer struct {
	port uint16
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minerWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minerWallet.PrivateKeyStr())
		log.Printf("public_key %v", minerWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minerWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := json.Marshal(bc)
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
	}
}

func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		transactions := bc.TransactionPool()
		m, _ := json.Marshal(struct {
			Transactions []*blockchain.Transaction `json:"transactions"`
			Length       int                       `json:"length"`
		}{
			Transactions: transactions,
			Length:       len(transactions),
		})
		io.WriteString(w, string(m[:]))
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var btr blockchain.TransactionRequest
		e := decoder.Decode(&btr)
		if e != nil {
			log.Printf("ERROR: %v", e)
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request")))
			return
		}
		if !btr.Validate() {
			log.Println("ERROR: missing fields")
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request: missing fields")))
			return
		}
		publicKey := utils.PublicKeyFromString(*btr.SenderPublicKey)
		signature := utils.SignatureFromString(*btr.Signature)
		bc := bcs.GetBlockchain()
		isCreated := bc.CreateTransaction(
			*btr.SenderBlockchainAddress,
			*btr.RecipientBlockchainAddress,
			*btr.Value,
			publicKey,
			signature)
		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isCreated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("Transaction failed")
		} else {
			w.WriteHeader(http.StatusCreated)
			m = utils.JsonStatus("Transaction created")
		}
		io.WriteString(w, string(m[:]))
	case http.MethodPut:
		decoder := json.NewDecoder(req.Body)
		var btr blockchain.TransactionRequest
		e := decoder.Decode(&btr)
		if e != nil {
			log.Printf("ERROR: %v", e)
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request")))
			return
		}
		if !btr.Validate() {
			log.Println("ERROR: missing fields")
			io.WriteString(w, string(utils.JsonStatus("Invalid transaction request: missing fields")))
			return
		}
		publicKey := utils.PublicKeyFromString(*btr.SenderPublicKey)
		signature := utils.SignatureFromString(*btr.Signature)
		bc := bcs.GetBlockchain()
		isUpdated := bc.AddTransaction(
			*btr.SenderBlockchainAddress,
			*btr.RecipientBlockchainAddress,
			*btr.Value,
			publicKey,
			signature)
		w.Header().Add("Content-Type", "application/json")
		var m []byte
		if !isUpdated {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("Transaction failed")
		} else {
			m = utils.JsonStatus("Transaction created")
		}
		io.WriteString(w, string(m[:]))
	case http.MethodDelete:
		bc := bcs.GetBlockchain()
		bc.ClearTransactionPool()
		io.WriteString(w, string(utils.JsonStatus("Transaction pool cleared")))
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Mine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		isMined := bc.Mining()

		var m []byte
		if !isMined {
			w.WriteHeader(http.StatusBadRequest)
			m = utils.JsonStatus("failed")
		} else {
			m = utils.JsonStatus("success")
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) StartMining(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		bc.StartMining()
		m := utils.JsonStatus("success")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Amount(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		blockchainAddress := req.URL.Query().Get("blockchain_address")
		amount := bcs.GetBlockchain().CalculateTotalAmount(blockchainAddress)

		ar := &blockchain.AmountResponse{Amount: amount}
		m, _ := json.Marshal(ar)

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Consensus(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		bc := bcs.GetBlockchain()
		replaced := bc.ResolveConflicts()

		w.Header().Add("Content-Type", "application/json")
		if replaced {
			io.WriteString(w, string(utils.JsonStatus("Successfully reached consensus")))
		} else {
			io.WriteString(w, string(utils.JsonStatus("Failed to reach consensus")))
		}
	default:
		log.Printf("ERROR: Invalid HTTP request method: %v", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (bcs *BlockchainServer) Run() {
	bcs.GetBlockchain().Run()

	http.HandleFunc("/", bcs.GetChain)
	http.HandleFunc("/transactions", bcs.Transactions)
	http.HandleFunc("/mine", bcs.Mine)
	http.HandleFunc("/mine/start", bcs.StartMining)
	http.HandleFunc("/amount", bcs.Amount)
	http.HandleFunc("/consensus", bcs.Consensus)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.port)), nil))
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}
