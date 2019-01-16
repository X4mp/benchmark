package request

import (
	"errors"
	"fmt"
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

type request struct {
	UUID    *uuid.UUID                  `json:"id"`
	Clnt    client.Client               `json:"client"`
	Int     *time.Duration              `json:"interval"`
	Spts    []spot.Spot                 `json:"spots"`
	Cits    []city.City                 `json:"cities"`
	Regs    []region.Region             `json:"regions"`
	Countrs []country.Country           `json:"countries"`
	Orgs    []organization.Organization `json:"organizations"`
}

func createRequest(
	id *uuid.UUID,
	clint client.Client,
	intrvl *time.Duration,
	spts []spot.Spot,
	cities []city.City,
	regions []region.Region,
	countries []country.Country,
	orgs []organization.Organization,
) (Request, error) {
	out := request{
		UUID:    id,
		Clnt:    clint,
		Int:     intrvl,
		Spts:    spts,
		Cits:    cities,
		Regs:    regions,
		Countrs: countries,
		Orgs:    orgs,
	}

	return &out, nil
}

func createRequestFromNormalized(normalized *normalizedRequest) (Request, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	cl, clErr := client.SDKFunc.CreateMetaData().Denormalize()(normalized.Client)
	if clErr != nil {
		return nil, clErr
	}

	spots := []spot.Spot{}
	if len(normalized.Spots) > 0 {
		for _, nSpot := range normalized.Spots {
			oneSpot, oneSpotErr := spot.SDKFunc.CreateMetaData().Denormalize()(nSpot)
			if oneSpotErr != nil {
				return nil, oneSpotErr
			}

			if spt, ok := oneSpot.(spot.Spot); ok {
				spots = append(spots, spt)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the spot list that is not a valid Spot instance", oneSpot.ID().String())
			return nil, errors.New(str)
		}
	}

	cities := []city.City{}
	if len(normalized.Cities) > 0 {
		for _, nCity := range normalized.Spots {
			oneCity, oneCityErr := city.SDKFunc.CreateMetaData().Denormalize()(nCity)
			if oneCityErr != nil {
				return nil, oneCityErr
			}

			if cit, ok := oneCity.(city.City); ok {
				cities = append(cities, cit)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the city list that is not a valid Spot instance", oneCity.ID().String())
			return nil, errors.New(str)
		}
	}

	regions := []region.Region{}
	if len(normalized.Regions) > 0 {
		for _, nRegion := range normalized.Regions {
			oneRegion, oneRegionErr := region.SDKFunc.CreateMetaData().Denormalize()(nRegion)
			if oneRegionErr != nil {
				return nil, oneRegionErr
			}

			if reg, ok := oneRegion.(region.Region); ok {
				regions = append(regions, reg)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the city list that is not a valid Spot instance", oneRegion.ID().String())
			return nil, errors.New(str)
		}
	}

	countries := []country.Country{}
	if len(normalized.Countries) > 0 {
		for _, nCountry := range normalized.Countries {
			oneCountry, oneCountryErr := country.SDKFunc.CreateMetaData().Denormalize()(nCountry)
			if oneCountryErr != nil {
				return nil, oneCountryErr
			}

			if countr, ok := oneCountry.(country.Country); ok {
				countries = append(countries, countr)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the country list that is not a valid Spot instance", oneCountry.ID().String())
			return nil, errors.New(str)
		}
	}

	orgs := []organization.Organization{}
	if len(normalized.Organizations) > 0 {
		for _, nOrganization := range normalized.Organizations {
			oneOrg, oneOrgErr := organization.SDKFunc.CreateMetaData().Denormalize()(nOrganization)
			if oneOrgErr != nil {
				return nil, oneOrgErr
			}

			if org, ok := oneOrg.(organization.Organization); ok {
				orgs = append(orgs, org)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the organization list that is not a valid Spot instance", oneOrg.ID().String())
			return nil, errors.New(str)
		}
	}

	if cli, ok := cl.(client.Client); ok {
		return createRequest(&id, cli, normalized.Interval, spots, cities, regions, countries, orgs)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", cl.ID().String())
	return nil, errors.New(str)

}

func createRequestFromStorable(storable *storableRequest, rep entity.Repository) (Request, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	clientID, clientIDErr := uuid.FromString(storable.ClientID)
	if clientIDErr != nil {
		return nil, clientIDErr
	}

	cl, clErr := rep.RetrieveByID(client.SDKFunc.CreateMetaData(), &clientID)
	if clErr != nil {
		return nil, clErr
	}

	spots := []spot.Spot{}
	if len(storable.SpotIDs) > 0 {
		for _, spotIDAsString := range storable.SpotIDs {

			spotID, spotIDErr := uuid.FromString(spotIDAsString)
			if spotIDErr != nil {
				return nil, spotIDErr
			}

			oneSpot, oneSpotErr := rep.RetrieveByID(spot.SDKFunc.CreateMetaData(), &spotID)
			if oneSpotErr != nil {
				return nil, oneSpotErr
			}

			if spt, ok := oneSpot.(spot.Spot); ok {
				spots = append(spots, spt)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the spot list that is not a valid Spot instance", oneSpot.ID().String())
			return nil, errors.New(str)
		}
	}

	cities := []city.City{}
	if len(storable.CityIDs) > 0 {
		for _, cityIDAsString := range storable.CityIDs {

			cityID, cityIDErr := uuid.FromString(cityIDAsString)
			if cityIDErr != nil {
				return nil, cityIDErr
			}

			oneCity, oneCityErr := rep.RetrieveByID(city.SDKFunc.CreateMetaData(), &cityID)
			if oneCityErr != nil {
				return nil, oneCityErr
			}

			if cit, ok := oneCity.(city.City); ok {
				cities = append(cities, cit)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the city list that is not a valid Spot instance", oneCity.ID().String())
			return nil, errors.New(str)
		}
	}

	regions := []region.Region{}
	if len(storable.RegionIDs) > 0 {
		for _, regionIDAsString := range storable.RegionIDs {

			regionID, regionIDErr := uuid.FromString(regionIDAsString)
			if regionIDErr != nil {
				return nil, regionIDErr
			}

			oneRegion, oneRegionErr := rep.RetrieveByID(region.SDKFunc.CreateMetaData(), &regionID)
			if oneRegionErr != nil {
				return nil, oneRegionErr
			}

			if reg, ok := oneRegion.(region.Region); ok {
				regions = append(regions, reg)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the city list that is not a valid Spot instance", oneRegion.ID().String())
			return nil, errors.New(str)
		}
	}

	countries := []country.Country{}
	if len(storable.CountryIDs) > 0 {
		for _, countryIDAsString := range storable.CountryIDs {

			countryID, countryIDErr := uuid.FromString(countryIDAsString)
			if countryIDErr != nil {
				return nil, countryIDErr
			}

			oneCountry, oneCountryErr := rep.RetrieveByID(country.SDKFunc.CreateMetaData(), &countryID)
			if oneCountryErr != nil {
				return nil, oneCountryErr
			}

			if countr, ok := oneCountry.(country.Country); ok {
				countries = append(countries, countr)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the country list that is not a valid Spot instance", oneCountry.ID().String())
			return nil, errors.New(str)
		}
	}

	orgs := []organization.Organization{}
	if len(storable.OrganizationIDs) > 0 {
		for _, orgIDAsString := range storable.OrganizationIDs {

			orgID, orgIDErr := uuid.FromString(orgIDAsString)
			if orgIDErr != nil {
				return nil, orgIDErr
			}

			oneOrg, oneOrgErr := rep.RetrieveByID(organization.SDKFunc.CreateMetaData(), &orgID)
			if oneOrgErr != nil {
				return nil, oneOrgErr
			}

			if org, ok := oneOrg.(organization.Organization); ok {
				orgs = append(orgs, org)
				continue
			}

			str := fmt.Sprintf("there is one entity (ID: %s) in the organization list that is not a valid Spot instance", oneOrg.ID().String())
			return nil, errors.New(str)
		}
	}

	if cli, ok := cl.(client.Client); ok {
		return createRequest(&id, cli, storable.Interval, spots, cities, regions, countries, orgs)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", cl.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *request) ID() *uuid.UUID {
	return obj.UUID
}

// Client returns the client
func (obj *request) Client() client.Client {
	return obj.Clnt
}

// HasInterval returns true if there is an interval, false otherwise
func (obj *request) HasInterval() bool {
	return obj.Int != nil
}

// Interval returns the interval if any
func (obj *request) Interval() *time.Duration {
	return obj.Int
}

// HasSpot returns true if there is a spot, false otherwise
func (obj *request) HasSpots() bool {
	return len(obj.Spts) > 0
}

// Spots returns the spots
func (obj *request) Spots() []spot.Spot {
	return obj.Spts
}

// HasCities returns true if there is cities, false otherwise
func (obj *request) HasCities() bool {
	return len(obj.Cits) > 0
}

// Cities returns the cities, if any
func (obj *request) Cities() []city.City {
	return obj.Cits
}

// HasRegions returns true if there is regions, false otherwise
func (obj *request) HasRegions() bool {
	return len(obj.Regs) > 0
}

// Regions returns the regions, if any
func (obj *request) Regions() []region.Region {
	return obj.Regs
}

// HasCountries returns true if there is countries, false otherwise
func (obj *request) HasCountries() bool {
	return len(obj.Countrs) > 0
}

// Countries returns the countries, if any
func (obj *request) Countries() []country.Country {
	return obj.Countrs
}

// HasOrganizations returns true if there is organizations, false otherwise
func (obj *request) HasOrganizations() bool {
	return len(obj.Orgs) > 0
}

// Organizations returns the organizations, if any
func (obj *request) Organizations() []organization.Organization {
	return obj.Orgs
}
