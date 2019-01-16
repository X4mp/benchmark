package customer

import (
	"errors"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnnetwork/benchmark/objects/report/server"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

type customer struct {
	UUID      *uuid.UUID        `json:"id"`
	Trsf      transfer.Transfer `json:"transfer"`
	Reps      []report.Report   `json:"report"`
	IPAddress net.IP            `json:"address"`
	Prt       int               `json:"port"`
}

func createCustomer(id *uuid.UUID, trsf transfer.Transfer, reps []report.Report) (Customer, error) {
	if len(reps) <= 0 {
		str := fmt.Sprintf("the customer must order at least 1 report, %d given", len(reps))
		return nil, errors.New(str)
	}

	out := customer{
		UUID: id,
		Trsf: trsf,
		Reps: reps,
	}

	return &out, nil
}

func createCustomerFromNormalized(normalized *normalizedCustomer) (Customer, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	trsfIns, trsfInsErr := transfer.SDKFunc.CreateMetaData().Denormalize()(normalized.Transfer)
	if trsfInsErr != nil {
		return nil, trsfInsErr
	}

	reps := []report.Report{}
	for _, oneNormalizedReport := range normalized.Reports {
		repIns, repInsErr := report.SDKFunc.CreateMetaData().Denormalize()(oneNormalizedReport)
		if repInsErr != nil {
			return nil, repInsErr
		}

		if rep, ok := repIns.(report.Report); ok {
			reps = append(reps, rep)
		}

		str := fmt.Sprintf("there is at least 1 entity (ID: %s) in the report list that is not a report instance", repIns.ID().String())
		return nil, errors.New(str)
	}

	if trsf, ok := trsfIns.(transfer.Transfer); ok {
		return createCustomer(&id, trsf, reps)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Transfer instance", trsfIns.ID().String())
	return nil, errors.New(str)
}

func createCustomerFromStorable(storable *storableCustomer, rep entity.Repository) (Customer, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	trsfID, trsfIDErr := uuid.FromString(storable.TransferID)
	if trsfIDErr != nil {
		return nil, trsfIDErr
	}

	reps := []report.Report{}
	repIDs := storable.ReportIDs
	for _, oneRepIDAsString := range repIDs {
		repID, repIDErr := uuid.FromString(oneRepIDAsString)
		if repIDErr != nil {
			return nil, repIDErr
		}

		repIns, repInsErr := rep.RetrieveByID(server.SDKFunc.CreateMetaData(), &repID)
		if repInsErr != nil {
			return nil, repInsErr
		}

		if rep, ok := repIns.(report.Report); ok {
			reps = append(reps, rep)
			continue
		}

		str := fmt.Sprintf("there is at least 1 entity (ID: %s) in the report list that is not a report instance", repIns.ID().String())
		return nil, errors.New(str)
	}

	trsfIns, trsfInsErr := rep.RetrieveByID(transfer.SDKFunc.CreateMetaData(), &trsfID)
	if trsfInsErr != nil {
		return nil, trsfInsErr
	}

	if trsf, ok := trsfIns.(transfer.Transfer); ok {
		return createCustomer(&id, trsf, reps)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Transfer instance", trsfIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *customer) ID() *uuid.UUID {
	return obj.UUID
}

// Transfer returns the transfer
func (obj *customer) Transfer() transfer.Transfer {
	return obj.Trsf
}

// Report returns the report
func (obj *customer) Reports() []report.Report {
	return obj.Reps
}

// IP returns the ip
func (obj *customer) IP() net.IP {
	return obj.IPAddress
}

// Port returns the port
func (obj *customer) Port() int {
	return obj.Prt
}
