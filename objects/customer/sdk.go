package customer

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

// Customer represents a customer paying for a report
type Customer interface {
	ID() *uuid.UUID
	Transfer() transfer.Transfer
	Report() report.Report
}

// Normalized represents a normalized customer
type Normalized interface {
}

// Repository represents a customer repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Customer, error)
	RetrieveByTransfer(trsf transfer.Transfer) (Customer, error)
	RetrieveSetByReport(rep report.Report, index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID       *uuid.UUID
	Transfer transfer.Transfer
	Report   report.Report
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Customer SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Customer
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Customer {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createCustomer(params.ID, params.Transfer, params.Report)
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
