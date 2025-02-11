package main

import (
	"evochain/constants"
	"log"
)

func init() {
	log.SetPrefix(constants.BLOCKCHAIN_NAME + ": ")
}

func main() {
	log.Println("Hello  world !")
}
