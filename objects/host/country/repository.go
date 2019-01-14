package country

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
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

// RetrieveByID retrieves a country by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Country, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if country, ok := ins.(Country); ok {
		return country, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByName retrieves a country by name
func (app *repository) RetrieveByName(name string) (Country, error) {
	keynames := []string{
		retrieveAllCountryKeyname(),
		retrieveCountryByNameKeyname(name),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if country, ok := ins.(Country); ok {
		return country, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSet retrieves a country set
func (app *repository) RetrieveSet(index int, amount int) (entity.PartialSet, error) {
	keyname := retrieveAllCountryKeyname()
	return app.entityRepository.RetrieveSetByKeyname(app.metaData, keyname, index, amount)
}
