package server

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
)

// Server represents a server
type Server interface {
	ID() *uuid.UUID
	Pledge() pledge.Pledge
	Host() host.Host
}

// Normalized represents a normalized server
type Normalized interface {
}

// Repository represents the client repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Server, error)
	RetrieveByHost(hst host.Host) (Server, error)
	RetrieveByPledge(pldge pledge.Pledge) (Server, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID     *uuid.UUID
	Pledge pledge.Pledge
	Host   host.Host
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Server SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Server
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Server {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createServer(params.ID, params.Pledge, params.Host)
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
