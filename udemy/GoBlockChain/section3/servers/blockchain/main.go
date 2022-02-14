package main

import (
	"flag"
	"fmt"
	"log"

	"udemy.com/goblockchain/section3/servers/blockchain/server"
)

func init() {
	log.SetPrefix("Blockchain Server:")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	fmt.Println("Blockchain Server Starting...")
	fmt.Println("PORT:", *port)
	app := server.NewBlockchainServer(uint16(*port))
	app.Run()
}
