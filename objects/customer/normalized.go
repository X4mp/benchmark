package customer

import (
	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

type normalizedCustomer struct {
	ID       string              `json:"id"`
	Transfer transfer.Normalized `json:"transfer"`
	Reports  []report.Normalized `json:"reports"`
}

func createNormalizedCustomer(ins Customer) (*normalizedCustomer, error) {
	trsf, trsfErr := transfer.SDKFunc.CreateMetaData().Normalize()(ins.Transfer())
	if trsfErr != nil {
		return nil, trsfErr
	}

	nReps := []report.Normalized{}
	reps := ins.Reports()
	for _, oneRep := range reps {
		rep, repErr := report.SDKFunc.CreateMetaData().Normalize()(oneRep)
		if repErr != nil {
			return nil, repErr
		}

		nReps = append(nReps, rep)
	}

	out := normalizedCustomer{
		ID:       ins.ID().String(),
		Transfer: trsf,
		Reports:  nReps,
	}

	return &out, nil
}
