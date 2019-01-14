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
	Rep       report.Report     `json:"report"`
	IPAddress net.IP            `json:"address"`
	Prt       int               `json:"port"`
}

func createCustomer(id *uuid.UUID, trsf transfer.Transfer, rep report.Report) (Customer, error) {
	out := customer{
		UUID: id,
		Trsf: trsf,
		Rep:  rep,
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

	repIns, repInsErr := report.SDKFunc.CreateMetaData().Denormalize()(normalized.Report)
	if repInsErr != nil {
		return nil, repInsErr
	}

	if trsf, ok := trsfIns.(transfer.Transfer); ok {
		if rep, ok := repIns.(report.Report); ok {
			return createCustomer(&id, trsf, rep)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid server Report instance", repIns.ID().String())
		return nil, errors.New(str)
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

	repID, repIDErr := uuid.FromString(storable.ReportID)
	if repIDErr != nil {
		return nil, repIDErr
	}

	trsfIns, trsfInsErr := rep.RetrieveByID(transfer.SDKFunc.CreateMetaData(), &trsfID)
	if trsfInsErr != nil {
		return nil, trsfInsErr
	}

	repIns, repInsErr := rep.RetrieveByID(server.SDKFunc.CreateMetaData(), &repID)
	if repInsErr != nil {
		return nil, repInsErr
	}

	if trsf, ok := trsfIns.(transfer.Transfer); ok {
		if rep, ok := repIns.(report.Report); ok {
			return createCustomer(&id, trsf, rep)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid server Report instance", repIns.ID().String())
		return nil, errors.New(str)
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
func (obj *customer) Report() report.Report {
	return obj.Rep
}

// IP returns the ip
func (obj *customer) IP() net.IP {
	return obj.IPAddress
}

// Port returns the port
func (obj *customer) Port() int {
	return obj.Prt
}
