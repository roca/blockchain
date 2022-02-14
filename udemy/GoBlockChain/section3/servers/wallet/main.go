package main

import (
	"flag"
	"fmt"
	"log"

	"udemy.com/goblockchain/section3/servers/wallet/server"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Blockchain Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway URL")
	flag.Parse()
	fmt.Println("Wallet Server Starting...")
	fmt.Println("PORT:", *port)
	fmt.Println("Backend GATEWAY:", *gateway)
	app := server.NewWallServer(uint16(*port), *gateway)
	app.Run()
}
