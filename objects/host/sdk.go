package host

import (
	"net"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Host represents a host
type Host interface {
	ID() *uuid.UUID
	IP() net.IP
	Hostname() string
	Longitude() int
	Latitude() int
	HasCountry() bool
	Country() country.Country
	HasRegion() bool
	Region() region.Region
	HasCity() bool
	City() city.City
	HasOrganization() bool
	Organization() organization.Organization
}

// Normalized represents a normalized host
type Normalized interface {
}

// Repository represents a host repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Host, error)
	RetrieveSetByIP(ip net.IP, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByCountry(count country.Country, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByRegion(reg region.Region, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByCity(cit city.City, index int, amount int) (entity.PartialSet, error)
	RetrieveSetByOrganization(org organization.Organization, index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID           *uuid.UUID
	IP           net.IP
	Hostname     string
	Longitude    int
	Latitude     int
	Country      country.Country
	Region       region.Region
	City         city.City
	Organization organization.Organization
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Host SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Host
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Host {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		if params.City != nil {
			out, outErr := createHostWithCity(
				params.ID,
				params.IP,
				params.Hostname,
				params.Longitude,
				params.Latitude,
				params.Organization,
				params.City,
			)

			if outErr != nil {
				panic(outErr)
			}

			return out
		}

		if params.Region != nil {
			out, outErr := createHostWithRegion(
				params.ID,
				params.IP,
				params.Hostname,
				params.Longitude,
				params.Latitude,
				params.Organization,
				params.Region,
			)

			if outErr != nil {
				panic(outErr)
			}

			return out
		}

		out, outErr := createHostWithCountry(
			params.ID,
			params.IP,
			params.Hostname,
			params.Longitude,
			params.Latitude,
			params.Organization,
			params.Country,
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
