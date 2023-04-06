package api

import (
	"net/http"
)

var Routes = []struct {
	Path    string
	Handler func() // function signature should match the expected handler function signature
	Method  string
}{
	{
		Path:    "/health",
		Handler: handleHealth,
		Method:  http.MethodGet,
	},
	{
		Path:    "/utils/address/{addr}/validate",
		Handler: handleAddressValidator,
		Method:  http.MethodGet,
	},
	{
		Path:    "/utils/keypair/random",
		Handler: handleCreateRandomFullKeypair,
		Method:  http.MethodGet,
	},
	{
		Path:    "/fee/stats",
		Handler: handleGetFeeStats,
		Method:  http.MethodGet,
	},
	{
		Path:    "/account/{account}/balance",
		Handler: handleBalance,
		Method:  http.MethodGet,
	},
}
