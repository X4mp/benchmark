package server

import (
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
)

type normalizedServer struct {
	ID     string            `json:"id"`
	Pledge pledge.Normalized `json:"pledge"`
	Host   host.Normalized   `json:"host"`
}

func createNormalizedServer(ins Server) (*normalizedServer, error) {
	pldge, pldgeErr := pledge.SDKFunc.CreateMetaData().Normalize()(ins.Pledge())
	if pldgeErr != nil {
		return nil, pldgeErr
	}

	hst, hstErr := host.SDKFunc.CreateMetaData().Normalize()(ins.Host())
	if hstErr != nil {
		return nil, hstErr
	}

	out := normalizedServer{
		ID:     ins.ID().String(),
		Pledge: pldge,
		Host:   hst,
	}

	return &out, nil
}
