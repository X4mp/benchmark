package report

import (
	"github.com/xmnnetwork/benchmark/objects/client"
)

type normalizedReport struct {
	ID     string            `json:"id"`
	Client client.Normalized `json:"client"`
}

func createNormalizedReport(ins Report) (*normalizedReport, error) {
	cl, clErr := client.SDKFunc.CreateMetaData().Normalize()(ins.Client())
	if clErr != nil {
		return nil, clErr
	}

	out := normalizedReport{
		ID:     ins.ID().String(),
		Client: cl,
	}

	return &out, nil
}
