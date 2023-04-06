package fee

import (
	"log"
	"strconv"

	"github.com/stellar/go/clients/horizonclient"
)

type Fee struct {
	high   string
	medium string
	low    string
}

func GetFees() (Fee, error) {
	log.Printf("[GetFees]")

	fs, e := horizonclient.DefaultPublicNetClient.FeeStats()

	if e != nil {
		return Fee{}, e
	}

	return Fee{
		high:   strconv.FormatInt(fs.FeeCharged.Max, 10),
		medium: strconv.FormatInt(fs.FeeCharged.Mode, 10),
		low:    strconv.FormatInt(fs.FeeCharged.Min, 10),
	}, nil
}
