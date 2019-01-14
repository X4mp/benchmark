package server

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
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

// RetrieveByID retrieves the server by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Server, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if serv, ok := ins.(Server); ok {
		return serv, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Pledge instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByHost retrieves the server by host
func (app *repository) RetrieveByHost(hst host.Host) (Server, error) {
	keynames := []string{
		retrieveAllServerKeyname(),
		retrieveServerByHostKeyname(hst),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if serv, ok := ins.(Server); ok {
		return serv, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Pledge instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByPledge retrieves the server by pledge
func (app *repository) RetrieveByPledge(pldge pledge.Pledge) (Server, error) {
	keynames := []string{
		retrieveAllServerKeyname(),
		retrieveServerByPledgeKeyname(pldge),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if serv, ok := ins.(Server); ok {
		return serv, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Pledge instance", ins.ID().String())
	return nil, errors.New(str)
}
