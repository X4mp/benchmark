package request

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Request represents a client request
type Request interface {
	ID() *uuid.UUID
	Client() client.Client
	HasInterval() bool
	Interval() *time.Duration
	HasSpots() bool
	Spots() []spot.Spot
	HasCities() bool
	Cities() []city.City
	HasRegions() bool
	Regions() []region.Region
	HasCountries() bool
	Countries() []country.Country
	HasOrganizations() bool
	Organizations() []organization.Organization
}

// Normalized represents a normalized request
type Normalized interface {
}

// Repository represents a request repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Request, error)
	RetrieveSetByClient(cl client.Client, index int, amount int) (entity.PartialSet, error)
	RetrieveSetBySpot(spt spot.Spot, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByCity(cit city.City, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByRegion(reg region.Region, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByCountry(count country.Country, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByOrganization(org organization.Organization, index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID            *uuid.UUID
	Client        client.Client
	Interval      *time.Duration
	Spots         []spot.Spot
	Cities        []city.City
	Regions       []region.Region
	Countries     []country.Country
	Organizations []organization.Organization
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Customer SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Request
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Request {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createRequest(
			params.ID,
			params.Client,
			params.Interval,
			params.Spots,
			params.Cities,
			params.Regions,
			params.Countries,
			params.Organizations,
		)

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
