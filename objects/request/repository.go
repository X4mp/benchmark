package request

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
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

// RetrieveByID retrieves a request by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Request, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if req, ok := ins.(Request); ok {
		return req, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Request instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSetByClient retrieves a request set by client
func (app *repository) RetrieveSetByClient(cl client.Client, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestByClientKeyname(cl),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetBySpot retrieves a request set by spot
func (app *repository) RetrieveSetBySpot(spt spot.Spot, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestBySpotKeyname(spt),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByCity retrieves a request set by city
func (app *repository) RetrieveSetByCity(cit city.City, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestByCityKeyname(cit),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByRegion retrieves a request set by region
func (app *repository) RetrieveSetByRegion(reg region.Region, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestByRegionKeyname(reg),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByCountry retrieves a request set by country
func (app *repository) RetrieveSetByCountry(count country.Country, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestByCountryKeyname(count),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByOrganization retrieves a request set by organization
func (app *repository) RetrieveSetByOrganization(org organization.Organization, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllRequestKeyname(),
		retrieveRequestByOrganizationKeyname(org),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}
