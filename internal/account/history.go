package account

import (
	"log"

	"github.com/stellar/go/clients/horizonclient"
)

type PaymentOp struct {
	Amount    string
	AssetCode string
	From      string
	To        string
}

type Order string

const (
	Ascending  Order = "asc"
	Descending Order = "desc"
)

type GetAccountPaymentsHistoryParams struct {
	Limit     uint
	Order     Order
	Page      string
	AccountId string
}

type AccountPaymentsHistory struct {
	limit    uint
	order    Order
	page     string
	Payments []PaymentOp
}

func GetAccountPaymentssHistory(req GetAccountPaymentsHistoryParams) (AccountPaymentsHistory, error) {
	log.Printf("[GetAccountPaymentsHistory] accountId: %s, limit: %d, order: %s, page: %d", r.AccountId, r.Limit, r.Order, r.Page)

	// This method implements the following endpoint:
	// GET https://horizon.stellar.org/accounts/{id}/payments{?cursor,limit,order}

	data, e := horizonclient.DefaultPublicNetClient.Payments(horizonclient.PaymentRequest{
		ForAccount: req.AccountId,
		Limit:      req.Limit,
		Order:      horizonclient.Order(req.Order),
		Cursor:     req.Page,
	})

	if e != nil {
		return AccountPaymentsHistory{}, e
	}

	payments := make([]PaymentOp, len(data.Embedded.Records))

	for i, p := range data.Embedded.Records {
		payments[i] = PaymentOp{
			Amount:    p.Amount,
			AssetCode: p.Asset.Code,
			From:      p.From,
			To:        p.To,
		}
	}

	return AccountPaymentsHistory{
		limit:    req.Limit,
		order:    req.Order,
		page:     data.Links.Next.Href,
		Payments: payments,
	}, nil
}
