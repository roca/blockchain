package main

import (
	"flag"
	"log"

	"udemy.com/goblockchain/section3/server/blockchain_server"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := blockchain_server.NewBlockchainServer(uint16(*port))
	app.Run()
}
