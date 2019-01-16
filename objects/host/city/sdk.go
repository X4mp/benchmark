package city

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// City represents a city
type City interface {
	ID() *uuid.UUID
	Name() string
	HasRegion() bool
	Region() region.Region
	HasCountry() bool
	Country() country.Country
}

// Normalized represents a normalized city
type Normalized interface {
}

// Repository represents a city repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (City, error)
	RetrieveByName(name string) (City, error)
	RetrieveSet(index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID      *uuid.UUID
	Name    string
	Region  region.Region
	Country country.Country
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the City SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) City
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) City {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		if params.Region != nil {
			out, outErr := createCityWithRegion(params.ID, params.Name, params.Region)
			if outErr != nil {
				panic(outErr)
			}

			return out
		}

		out, outErr := createCityWithCountry(params.ID, params.Name, params.Country)
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
