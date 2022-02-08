package main

import (
	"flag"
	"log"

	"udemy.com/goblockchain/section3/servers/blockchain/server"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := server.NewBlockchainServer(uint16(*port))
	app.Run()
}
