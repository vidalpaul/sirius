package account

import (
	"log"

	"github.com/stellar/go/clients/horizonclient"
)

type Balance struct {
	AssetCode      string
	Balance        string
	MinimumReserve string
}

func GetBalance(accountId string, assetCode string) (Balance, error) {
	log.Printf("[GetBalance]")

	a, e := horizonclient.DefaultPublicNetClient.AccountDetail(horizonclient.AccountRequest{AccountID: accountId})

	if e != nil {
		return Balance{}, e
	}

	balance := Balance{AssetCode: assetCode}

	for _, b := range a.Balances {
		if b.Asset.Code == assetCode {
			balance.Balance = b.Balance
			balance.MinimumReserve = b.Limit
		}
	}

	return balance, nil
}
