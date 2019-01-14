package client

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

type client struct {
	UUID      *uuid.UUID    `json:"id"`
	PaidToWal wallet.Wallet `json:"paid_to"`
	Hst       host.Host     `json:"host"`
}

func createClient(id *uuid.UUID, paidTo wallet.Wallet, hst host.Host) (Client, error) {
	out := client{
		UUID:      id,
		PaidToWal: paidTo,
		Hst:       hst,
	}

	return &out, nil
}

func createClientFromNormalized(normalized *normalizedClient) (Client, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	walIns, walInsErr := wallet.SDKFunc.CreateMetaData().Denormalize()(normalized.PaidTo)
	if walInsErr != nil {
		return nil, walInsErr
	}

	hstIns, hstInsErr := host.SDKFunc.CreateMetaData().Denormalize()(normalized.Host)
	if hstInsErr != nil {
		return nil, hstInsErr
	}

	if wal, ok := walIns.(wallet.Wallet); ok {
		if hst, ok := hstIns.(host.Host); ok {
			return createClient(&id, wal, hst)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", hstIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Wallet instance", walIns.ID().String())
	return nil, errors.New(str)
}

func createClientFromStorable(storable *storableClient, rep entity.Repository) (Client, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	walID, walIDErr := uuid.FromString(storable.PaidToWalletID)
	if walIDErr != nil {
		return nil, walIDErr
	}

	hostID, hostIDErr := uuid.FromString(storable.HostID)
	if hostIDErr != nil {
		return nil, hostIDErr
	}

	walIns, walInsErr := rep.RetrieveByID(wallet.SDKFunc.CreateMetaData(), &walID)
	if walInsErr != nil {
		return nil, walInsErr
	}

	hstIns, hstInsErr := rep.RetrieveByID(host.SDKFunc.CreateMetaData(), &hostID)
	if hstInsErr != nil {
		return nil, hstInsErr
	}

	if wal, ok := walIns.(wallet.Wallet); ok {
		if hst, ok := hstIns.(host.Host); ok {
			return createClient(&id, wal, hst)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", hstIns.ID().String())
		return nil, errors.New(str)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Wallet instance", walIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *client) ID() *uuid.UUID {
	return obj.UUID
}

// PaidTo returns the paidTo wallet
func (obj *client) PaidTo() wallet.Wallet {
	return obj.PaidToWal
}

// Host returns the host
func (obj *client) Host() host.Host {
	return obj.Hst
}
