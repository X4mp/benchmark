package server

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
)

type server struct {
	UUID  *uuid.UUID    `json:"id"`
	Pldge pledge.Pledge `json:"pledge"`
	Hst   host.Host     `json:"host"`
}

func createServer(id *uuid.UUID, pldge pledge.Pledge, hst host.Host) (Server, error) {
	out := server{
		UUID:  id,
		Pldge: pldge,
		Hst:   hst,
	}

	return &out, nil
}

func createServerFromNormalized(normalized *normalizedServer) (Server, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	pldgeIns, pldgeInsErr := pledge.SDKFunc.CreateMetaData().Denormalize()(normalized.Pledge)
	if pldgeInsErr != nil {
		return nil, pldgeInsErr
	}

	hstIns, hstInsErr := host.SDKFunc.CreateMetaData().Denormalize()(normalized.Host)
	if hstInsErr != nil {
		return nil, hstInsErr
	}

	if pldge, ok := pldgeIns.(pledge.Pledge); ok {
		if hst, ok := hstIns.(host.Host); ok {
			return createServer(&id, pldge, hst)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", hstIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Pledge instance", pldgeIns.ID().String())
	return nil, errors.New(str)
}

func createServerFromStorable(storable *storableServer, rep entity.Repository) (Server, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	pldgeID, pldgeIDErr := uuid.FromString(storable.PledgeID)
	if pldgeIDErr != nil {
		return nil, pldgeIDErr
	}

	hstID, hstIDErr := uuid.FromString(storable.HostID)
	if hstIDErr != nil {
		return nil, hstIDErr
	}

	pldgeIns, pldgeInsErr := rep.RetrieveByID(pledge.SDKFunc.CreateMetaData(), &pldgeID)
	if pldgeInsErr != nil {
		return nil, pldgeInsErr
	}

	hstIns, hstInsErr := rep.RetrieveByID(host.SDKFunc.CreateMetaData(), &hstID)
	if hstInsErr != nil {
		return nil, hstInsErr
	}

	if pldge, ok := pldgeIns.(pledge.Pledge); ok {
		if hst, ok := hstIns.(host.Host); ok {
			return createServer(&id, pldge, hst)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", hstIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Pledge instance", pldgeIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *server) ID() *uuid.UUID {
	return obj.UUID
}

// Pledge returns the pledge
func (obj *server) Pledge() pledge.Pledge {
	return obj.Pldge
}

// Host returns the host
func (obj *server) Host() host.Host {
	return obj.Hst
}
