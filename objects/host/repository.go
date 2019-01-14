package host

import (
	"errors"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type repository struct {
	metaData         entity.MetaData
	entityRepository entity.Repository
}

func createRepository(metaData entity.MetaData, entityRepository entity.Repository) Repository {
	out := repository{
		metaData:         metaData,
		entityRepository: entityRepository,
	}

	return &out
}

// RetrieveByID retrieves the host by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Host, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if hst, ok := ins.(Host); ok {
		return hst, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSetByIP retrieves host set by ip
func (app *repository) RetrieveSetByIP(ip net.IP, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllHostKeyname(),
		retrieveHostByIPKeyname(ip),
	}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByCountry retrieves host set by country
func (app *repository) RetrieveSetByCountry(count country.Country, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllHostKeyname(),
		retrieveHostByCountryKeyname(count),
	}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByRegion retrieves host set by region
func (app *repository) RetrieveSetByRegion(reg region.Region, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllHostKeyname(),
		retrieveHostByRegionKeyname(reg),
	}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByCity retrieves host set by city
func (app *repository) RetrieveSetByCity(cit city.City, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllHostKeyname(),
		retrieveHostByCityKeyname(cit),
	}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByOrganization retrieves host set by organization
func (app *repository) RetrieveSetByOrganization(org organization.Organization, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllHostKeyname(),
		retrieveHostByOrganizationKeyname(org),
	}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}
