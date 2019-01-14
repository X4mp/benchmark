package city

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

// RetrieveByID retrieves a city by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (City, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if city, ok := ins.(City); ok {
		return city, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByName retrieves a city by name
func (app *repository) RetrieveByName(name string) (City, error) {
	keynames := []string{
		retrieveAllCityKeyname(),
		retrieveCityByNameKeyname(name),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if city, ok := ins.(City); ok {
		return city, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSet retrieves a city set
func (app *repository) RetrieveSet(index int, amount int) (entity.PartialSet, error) {
	keyname := retrieveAllCityKeyname()
	return app.entityRepository.RetrieveSetByKeyname(app.metaData, keyname, index, amount)
}
