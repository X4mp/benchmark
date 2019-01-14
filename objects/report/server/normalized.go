package server

import (
	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
)

type normalizedReport struct {
	ID            string          `json:"id"`
	Report        rep.Normalized  `json:"report"`
	Server        serv.Normalized `json:"server"`
	EncryptedData string          `json:"encrypted_data"`
}

func createNormalizedReport(ins Report) (*normalizedReport, error) {
	reps, repsErr := rep.SDKFunc.CreateMetaData().Normalize()(ins.Report())
	if repsErr != nil {
		return nil, repsErr
	}

	servs, servsErr := serv.SDKFunc.CreateMetaData().Normalize()(ins.Server())
	if servsErr != nil {
		return nil, servsErr
	}

	out := normalizedReport{
		ID:            ins.ID().String(),
		Report:        reps,
		Server:        servs,
		EncryptedData: ins.EncryptedData(),
	}

	return &out, nil
}
