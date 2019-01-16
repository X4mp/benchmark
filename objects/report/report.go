package report

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/request"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type report struct {
	UUID *uuid.UUID      `json:"id"`
	Req  request.Request `json:"request"`
}

func createReport(id *uuid.UUID, req request.Request) (Report, error) {
	out := report{
		UUID: id,
		Req:  req,
	}

	return &out, nil
}

func createReportFromNormalized(normalized *normalizedReport) (Report, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	reqIns, reqInsErr := request.SDKFunc.CreateMetaData().Denormalize()(normalized.Request)
	if reqInsErr != nil {
		return nil, reqInsErr
	}

	if req, ok := reqIns.(request.Request); ok {
		return createReport(&id, req)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Request instance", reqIns.ID().String())
	return nil, errors.New(str)
}

func createReportFromStorable(storable *storableReport, rep entity.Repository) (Report, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	reqID, reqIDErr := uuid.FromString(storable.RequestID)
	if reqIDErr != nil {
		return nil, reqIDErr
	}

	reqIns, reqInsErr := rep.RetrieveByID(request.SDKFunc.CreateMetaData(), &reqID)
	if reqInsErr != nil {
		return nil, reqInsErr
	}

	if req, ok := reqIns.(request.Request); ok {
		return createReport(&id, req)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Request instance", reqIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *report) ID() *uuid.UUID {
	return obj.UUID
}

// Request returns the request
func (obj *report) Request() request.Request {
	return obj.Req
}
