package main

import (
	"fmt"

	"github.com/ArthurPaivaT/Porletu/mongohandler"
	"github.com/ArthurPaivaT/Porletu/server"
)

func main() {
	fmt.Println("Starting Server...")

	serverErrChan := make(chan error)

	go server.Start(serverErrChan)

	go mongohandler.Connect(serverErrChan)

	serverErr := <-serverErrChan
	fmt.Println(fmt.Errorf("Error Serving: %w", serverErr))

}
