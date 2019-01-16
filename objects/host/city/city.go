package city

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type city struct {
	UUID  *uuid.UUID      `json:"id"`
	Nme   string          `json:"name"`
	Reg   region.Region   `json:"region"`
	Count country.Country `json:"country"`
}

func createCityWithRegion(id *uuid.UUID, name string, reg region.Region) (City, error) {
	out := city{
		UUID:  id,
		Nme:   name,
		Reg:   reg,
		Count: nil,
	}

	return &out, nil
}

func createCityWithCountry(id *uuid.UUID, name string, count country.Country) (City, error) {
	out := city{
		UUID:  id,
		Nme:   name,
		Reg:   nil,
		Count: count,
	}

	return &out, nil
}

func createCityFromNormalized(normalized *normalizedCity) (City, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	if normalized.Country != nil {
		countryIns, countryInsErr := country.SDKFunc.CreateMetaData().Denormalize()(normalized.Country)
		if countryInsErr != nil {
			return nil, countryInsErr
		}

		if countr, ok := countryIns.(country.Country); ok {
			return createCityWithCountry(&id, normalized.Name, countr)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", countryIns.ID().String())
		return nil, errors.New(str)
	}

	regionIns, regionInsErr := region.SDKFunc.CreateMetaData().Denormalize()(normalized.Region)
	if regionInsErr != nil {
		return nil, regionInsErr
	}

	if reg, ok := regionIns.(region.Region); ok {
		return createCityWithRegion(&id, normalized.Name, reg)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regionIns.ID().String())
	return nil, errors.New(str)
}

func createCityFromStorable(storable *storableCity, rep entity.Repository) (City, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	if storable.CountryID != "" {
		countryID, countryIDErr := uuid.FromString(storable.CountryID)
		if countryIDErr != nil {
			return nil, countryIDErr
		}

		countryIns, countryInsErr := rep.RetrieveByID(country.SDKFunc.CreateMetaData(), &countryID)
		if countryInsErr != nil {
			return nil, countryInsErr
		}

		if countr, ok := countryIns.(country.Country); ok {
			return createCityWithCountry(&id, storable.Name, countr)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", countryIns.ID().String())
		return nil, errors.New(str)
	}

	regionID, regionIDErr := uuid.FromString(storable.RegionID)
	if regionIDErr != nil {
		return nil, regionIDErr
	}

	regionIns, regionInsErr := rep.RetrieveByID(region.SDKFunc.CreateMetaData(), &regionID)
	if regionInsErr != nil {
		return nil, regionInsErr
	}

	if reg, ok := regionIns.(region.Region); ok {
		return createCityWithRegion(&id, storable.Name, reg)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regionIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *city) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *city) Name() string {
	return obj.Nme
}

// HasRegion returns true if there is a region, false otherwise
func (obj *city) HasRegion() bool {
	return obj.Reg != nil
}

// Region returns the region
func (obj *city) Region() region.Region {
	return obj.Reg
}

// HasCountry returns true if there is country, false otherwise
func (obj *city) HasCountry() bool {
	return obj.Count != nil
}

// Country returns the country
func (obj *city) Country() country.Country {
	return obj.Count
}
