package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sirius/internal"
	"sirius/internal/api"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	server := api.NewServer(internal.GetAsInt("PORT", 8080))
	server.MustStart()
	defer server.Stop()

	server.AddRoute("/health", handleHealth(), http.MethodGet)

	server.AddRoute("/utils/address/{addr}/validate", handleAddressValidator(), http.MethodGet)

	/* server.AddRoute("/account/:account/balance", handleBalance(id), http.MethodGet)
	server.AddRoute("/account/:account/history", handleHistory(id), http.MethodGet)
	server.AddRoute("/tx/:tx/status", handleTxStatus(hash), http.MethodGet)
	*/
	// Wait for CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// We block here until a CTRL-C / SigInt is received
	// Once received, we exit and the server is cleaned up
	<-sigChan
}
