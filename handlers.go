package main

import (
	"net/http"
	"sirius/internal/api"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/strkey"
)

func handleHealth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ms := make(map[string]interface{})
		ms["status"] = "ok"
		api.JSONMessage(w, http.StatusOK, ms)
	})
}

func handleAddressValidator() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := api.GetParam(r, "addr")

		v := strkey.IsValidEd25519PublicKey(addr)

		ms := make(map[string]interface{})
		ms["valid"] = v

		api.JSONMessage(w, http.StatusOK, ms)
	})
}

func handleCreateRandomFullKeypair() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kp, e := keypair.Random()
		ms := make(map[string]interface{})
		if e != nil {
			ms["error"] = e.Error()
			api.JSONMessage(w, http.StatusInternalServerError, ms)
			return
		}

		ms["pk"] = kp.Address()
		ms["sk"] = kp.Seed()

		api.JSONMessage(w, http.StatusOK, ms)
	})
}

func handleGetFeeStats() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs, e := horizonclient.DefaultPublicNetClient.FeeStats()
		ms := make(map[string]interface{})
		if e != nil {
			ms["error"] = e.Error()
			api.JSONMessage(w, http.StatusInternalServerError, ms)
			return
		}

		ms["fee_stats"] = fs

		api.JSONMessage(w, http.StatusOK, ms)
	})
}
