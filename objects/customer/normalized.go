package customer

import (
	"github.com/xmnnetwork/benchmark/objects/report/server"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

type normalizedCustomer struct {
	ID       string              `json:"id"`
	Transfer transfer.Normalized `json:"transfer"`
	Report   server.Normalized   `json:"report"`
}

func createNormalizedCustomer(ins Customer) (*normalizedCustomer, error) {
	trsf, trsfErr := transfer.SDKFunc.CreateMetaData().Normalize()(ins.Transfer())
	if trsfErr != nil {
		return nil, trsfErr
	}

	rep, repErr := server.SDKFunc.CreateMetaData().Normalize()(ins.Report())
	if repErr != nil {
		return nil, repErr
	}

	out := normalizedCustomer{
		ID:       ins.ID().String(),
		Transfer: trsf,
		Report:   rep,
	}

	return &out, nil
}
