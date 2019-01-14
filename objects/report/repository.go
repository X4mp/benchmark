package report

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/client"
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

// RetrieveByID retrieves the report by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Report, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if rep, ok := ins.(Report); ok {
		return rep, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Information instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSetByClient retrieves a report set by client
func (app *repository) RetrieveSetByClient(cl client.Client, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}
