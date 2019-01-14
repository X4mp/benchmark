package client

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

// Client represents a client
type Client interface {
	ID() *uuid.UUID
	PaidTo() wallet.Wallet
	Host() host.Host
}

// Normalized represents a normalized client
type Normalized interface {
}

// Repository represents the client repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Client, error)
	RetrieveByPaidToWallet(paidTo wallet.Wallet) (Client, error)
	RetrieveByHost(hst host.Host) (Client, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID     *uuid.UUID
	PaidTo wallet.Wallet
	Host   host.Host
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Client SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Client
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Client {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createClient(params.ID, params.PaidTo, params.Host)

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
