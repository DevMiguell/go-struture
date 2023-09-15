package main

import (
	"flag"
	"fmt"
	"log"
	"v1/api"
	"v1/storage"
)

func main() {
	linstenAddr := flag.String("linstenAddr", ":3000", "address of the server")
	flag.Parse()

	storage := storage.NewMemoryStorage()

	server := api.NewServer(*linstenAddr, storage)
	fmt.Println("Server running on port: ", *linstenAddr)
	log.Fatal(server.Start())
}
