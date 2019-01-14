package report

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type report struct {
	UUID *uuid.UUID    `json:"id"`
	Clnt client.Client `json:"client"`
}

func createReport(id *uuid.UUID, clnt client.Client) (Report, error) {
	out := report{
		UUID: id,
		Clnt: clnt,
	}

	return &out, nil
}

func createReportFromNormalized(normalized *normalizedReport) (Report, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	clIns, clInsErr := client.SDKFunc.CreateMetaData().Denormalize()(normalized.Client)
	if clInsErr != nil {
		return nil, clInsErr
	}

	if cl, ok := clIns.(client.Client); ok {
		return createReport(&id, cl)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", clIns.ID().String())
	return nil, errors.New(str)
}

func createReportFromStorable(storable *storableReport, rep entity.Repository) (Report, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	clientID, clientIDErr := uuid.FromString(storable.ClientID)
	if clientIDErr != nil {
		return nil, clientIDErr
	}

	clIns, clInsErr := rep.RetrieveByID(client.SDKFunc.CreateMetaData(), &clientID)
	if clInsErr != nil {
		return nil, clInsErr
	}

	if cl, ok := clIns.(client.Client); ok {
		return createReport(&id, cl)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", clIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *report) ID() *uuid.UUID {
	return obj.UUID
}

// Client returns the client
func (obj *report) Client() client.Client {
	return obj.Clnt
}
