package client

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
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

// RetrieveByID retrieves the client by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Client, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if clint, ok := ins.(Client); ok {
		return clint, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByPaidToWallet retrieves the client by paidTo wallet
func (app *repository) RetrieveByPaidToWallet(paidTo wallet.Wallet) (Client, error) {
	keynames := []string{
		retrieveAllClientKeyname(),
		retrieveClientByPaidToWalletKeyname(paidTo),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if clint, ok := ins.(Client); ok {
		return clint, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByHost retrieves the client by host
func (app *repository) RetrieveByHost(hst host.Host) (Client, error) {
	keynames := []string{
		retrieveAllClientKeyname(),
		retrieveClientByHostKeyname(hst),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if clint, ok := ins.(Client); ok {
		return clint, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", ins.ID().String())
	return nil, errors.New(str)
}
