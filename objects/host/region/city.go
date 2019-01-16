package region

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type region struct {
	UUID  *uuid.UUID      `json:"id"`
	Nme   string          `json:"name"`
	Count country.Country `json:"country"`
}

func createRegion(id *uuid.UUID, name string, count country.Country) (Region, error) {
	out := region{
		UUID:  id,
		Nme:   name,
		Count: count,
	}

	return &out, nil
}

func createRegionFromNormalized(normalized *normalizedRegion) (Region, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	countryIns, countryInsErr := country.SDKFunc.CreateMetaData().Denormalize()(normalized.Country)
	if countryInsErr != nil {
		return nil, countryInsErr
	}

	if count, ok := countryIns.(country.Country); ok {
		return createRegion(&id, normalized.Name, count)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", countryIns.ID().String())
	return nil, errors.New(str)
}

func createRegionFromStorable(storable *storableRegion, rep entity.Repository) (Region, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	countryID, countryIDErr := uuid.FromString(storable.CountryID)
	if countryIDErr != nil {
		return nil, countryIDErr
	}

	countryIns, countryInsErr := rep.RetrieveByID(country.SDKFunc.CreateMetaData(), &countryID)
	if countryInsErr != nil {
		return nil, countryInsErr
	}

	if count, ok := countryIns.(country.Country); ok {
		return createRegion(&id, storable.Name, count)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", countryIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *region) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *region) Name() string {
	return obj.Nme
}

// Country returns the country
func (obj *region) Country() country.Country {
	return obj.Count
}
