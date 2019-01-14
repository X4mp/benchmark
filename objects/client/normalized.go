package client

import (
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

type normalizedClient struct {
	ID     string            `json:"id"`
	PaidTo wallet.Normalized `json:"paid_to"`
	Host   host.Normalized   `json:"host"`
}

func createNormalizedClient(ins Client) (*normalizedClient, error) {
	wal, walErr := wallet.SDKFunc.CreateMetaData().Normalize()(ins.PaidTo())
	if walErr != nil {
		return nil, walErr
	}

	hst, hstErr := host.SDKFunc.CreateMetaData().Normalize()(ins.Host())
	if hstErr != nil {
		return nil, hstErr
	}

	out := normalizedClient{
		ID:     ins.ID().String(),
		PaidTo: wal,
		Host:   hst,
	}

	return &out, nil
}
