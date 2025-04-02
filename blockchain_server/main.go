package main

import (
	"blockchain/blockchain_server" // Ensure this matches the correct module path
	"flag"
	"log"
)

func init() {
	log.SetPrefix("BlockChain : ")
}

func main() {
	port := flag.Uint("port", 6000, "TCP Port Number for Blockchain Server")
	flag.Parse()

	app := blockchain_server.NewBlockchainServer(uint16(*port)) // Use the correct package
	app.Run()
}
