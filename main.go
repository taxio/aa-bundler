package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
)

const port = 8080

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := rpc.NewServer()
	if err := server.RegisterName("eth", &Bundler{}); err != nil {
		return err
	}
	if err := server.RegisterName("debug_bundler", &Debug{}); err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", port), server)
}
