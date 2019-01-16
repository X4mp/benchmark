package report

import (
	"github.com/xmnnetwork/benchmark/objects/request"
)

type normalizedReport struct {
	ID      string             `json:"id"`
	Request request.Normalized `json:"request"`
}

func createNormalizedReport(ins Report) (*normalizedReport, error) {
	req, reqErr := request.SDKFunc.CreateMetaData().Normalize()(ins.Request())
	if reqErr != nil {
		return nil, reqErr
	}

	out := normalizedReport{
		ID:      ins.ID().String(),
		Request: req,
	}

	return &out, nil
}
