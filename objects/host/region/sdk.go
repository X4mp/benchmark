package region

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Region represents a region
type Region interface {
	ID() *uuid.UUID
	Name() string
	Country() country.Country
}

// Normalized represents a normalized region
type Normalized interface {
}

// Repository represents a region repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Region, error)
	RetrieveByName(name string) (Region, error)
	RetrieveSet(index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID      *uuid.UUID
	Name    string
	Country country.Country
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Region SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Region
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Region {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createRegion(params.ID, params.Name, params.Country)
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
