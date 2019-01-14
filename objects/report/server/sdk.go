package server

import (
	uuid "github.com/satori/go.uuid"
	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Report represents a report made by a server
type Report interface {
	ID() *uuid.UUID
	Report() rep.Report
	Server() serv.Server
	EncryptedData() string
}

// Normalized represents a normalized host
type Normalized interface {
}

// Repository represents the report repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Report, error)
	RetrieveSetByReport(repo rep.Report, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByServer(serve serv.Server, index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID            *uuid.UUID
	Report        rep.Report
	Server        serv.Server
	EncryptedData string
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

		out, outErr := createReport(params.ID, params.Report, params.Server, params.EncryptedData)
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
