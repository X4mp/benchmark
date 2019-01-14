package region

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

// RetrieveByID retrieves a region by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Region, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if region, ok := ins.(Region); ok {
		return region, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByName retrieves a region by name
func (app *repository) RetrieveByName(name string) (Region, error) {
	keynames := []string{
		retrieveAllRegionKeyname(),
		retrieveRegionByNameKeyname(name),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if region, ok := ins.(Region); ok {
		return region, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSet retrieves a region set
func (app *repository) RetrieveSet(index int, amount int) (entity.PartialSet, error) {
	keyname := retrieveAllRegionKeyname()
	return app.entityRepository.RetrieveSetByKeyname(app.metaData, keyname, index, amount)
}
