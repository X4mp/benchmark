package organization

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

// RetrieveByID retrieves a organization by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Organization, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if organization, ok := ins.(Organization); ok {
		return organization, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByName retrieves a organization by name
func (app *repository) RetrieveByName(name string) (Organization, error) {
	keynames := []string{
		retrieveAllOrganizationKeyname(),
		retrieveOrganizationByNameKeyname(name),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if organization, ok := ins.(Organization); ok {
		return organization, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSet retrieves a organization set
func (app *repository) RetrieveSet(index int, amount int) (entity.PartialSet, error) {
	keyname := retrieveAllOrganizationKeyname()
	return app.entityRepository.RetrieveSetByKeyname(app.metaData, keyname, index, amount)
}
