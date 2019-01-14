package server

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type report struct {
	UUID *uuid.UUID  `json:"id"`
	Rep  rep.Report  `json:"report"`
	Serv serv.Server `json:"server"`
	Dat  string      `json:"encrypted_data"`
}

func createReport(id *uuid.UUID, rep rep.Report, serv serv.Server, data string) (Report, error) {
	out := report{
		UUID: id,
		Rep:  rep,
		Serv: serv,
		Dat:  data,
	}

	return &out, nil
}

func createReportFromNormalized(normalized *normalizedReport) (Report, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	repIns, repInsErr := rep.SDKFunc.CreateMetaData().Denormalize()(normalized.Report)
	if repInsErr != nil {
		return nil, repInsErr
	}

	servIns, servInsErr := serv.SDKFunc.CreateMetaData().Denormalize()(normalized.Server)
	if servInsErr != nil {
		return nil, servInsErr
	}

	if reps, ok := repIns.(rep.Report); ok {
		if servs, ok := servIns.(serv.Server); ok {
			return createReport(&id, reps, servs, normalized.EncryptedData)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Server instance", servIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Report instance", repIns.ID().String())
	return nil, errors.New(str)
}

func createReportFromStorable(storable *storableReport, entityRep entity.Repository) (Report, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	repID, repIDErr := uuid.FromString(storable.ReportID)
	if repIDErr != nil {
		return nil, repIDErr
	}

	servID, servIDErr := uuid.FromString(storable.ServerID)
	if servIDErr != nil {
		return nil, servIDErr
	}

	repIns, repInsErr := entityRep.RetrieveByID(rep.SDKFunc.CreateMetaData(), &repID)
	if repInsErr != nil {
		return nil, repInsErr
	}

	servIns, servInsErr := entityRep.RetrieveByID(serv.SDKFunc.CreateMetaData(), &servID)
	if servInsErr != nil {
		return nil, servInsErr
	}

	if reps, ok := repIns.(rep.Report); ok {
		if servs, ok := servIns.(serv.Server); ok {
			return createReport(&id, reps, servs, storable.EncryptedData)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Server instance", servIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Report instance", repIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *report) ID() *uuid.UUID {
	return obj.UUID
}

// Report returns the report
func (obj *report) Report() rep.Report {
	return obj.Rep
}

// Server returns the server
func (obj *report) Server() serv.Server {
	return obj.Serv
}

// Data returns the data
func (obj *report) EncryptedData() string {
	return obj.Dat
}
