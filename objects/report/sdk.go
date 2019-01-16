package report

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/request"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Report represents a report
type Report interface {
	ID() *uuid.UUID
	Request() request.Request
}

// Normalized represents a normalized report
type Normalized interface {
}

// Repository represents a report instance
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Report, error)
	RetrieveSetByRequest(req request.Request, index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID      *uuid.UUID
	Request request.Request
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Report SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Report
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Report {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createReport(params.ID, params.Request)
		if outErr != nil {
			panic(outErr)
		}

		return out
	},
	CreateMetaData: func() entity.MetaData {
		return createMetaData()
	},
	CreateRepresentation: func() entity.Representation {
		return representation()
	},
	CreateRepository: func(params CreateRepositoryParams) Repository {
		metaData := createMetaData()
		out := createRepository(metaData, params.EntityRepository)
		return out
	},
}
