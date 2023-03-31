package main

import (
	"net/http"
	"sirius/internal/api"
	"strconv"

	"github.com/stellar/go/strkey"
)

func handleHealth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.JSONMessage(w, http.StatusOK, "OK")
	})
}

func handleAddressValidator() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := api.GetParam(r, "addr")

		v := strkey.IsValidEd25519PublicKey(addr)

		api.JSONMessage(w, http.StatusOK, strconv.FormatBool(v))
	})
}
