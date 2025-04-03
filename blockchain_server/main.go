package main

import (
	   
	"flag"
	"log"

)

func init() {
	log.SetPrefix("BlockChain : ")
}

func main() {
	port := flag.Uint("port", 6000, "TCP Port Number for Blockchain Server")
	flag.Parse()

	app := NewBlockchainServer(uint16(*port)) // Use the correct package
	app.Run()
}
