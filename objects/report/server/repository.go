package server

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
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

	if repo, ok := ins.(Report); ok {
		return repo, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Report instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSetByReport retrieves the report set by report
func (app *repository) RetrieveSetByReport(repo rep.Report, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}

// RetrieveSetByServer retrieves the report set by server
func (app *repository) RetrieveSetByServer(serve serv.Server, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{}
	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}
