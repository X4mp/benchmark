package client

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllClientKeyname() string {
	return "clients"
}

func retrieveClientByPaidToWalletKeyname(paidTo wallet.Wallet) string {
	base := retrieveAllClientKeyname()
	return fmt.Sprintf("%s:by_paid_to_wallet_id:%s", base, paidTo.ID().String())
}

func retrieveClientByHostKeyname(hst host.Host) string {
	base := retrieveAllClientKeyname()
	return fmt.Sprintf("%s:by_host_id:%s", base, hst.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Client",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableClient); ok {
				return createClientFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedClient)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createClientFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Client instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if clnt, ok := ins.(Client); ok {
				out, outErr := createNormalizedClient(clnt)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Client instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedClient); ok {
				return createClientFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Client instance")
		},
		EmptyStorable:   new(storableClient),
		EmptyNormalized: new(normalizedClient),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if clnt, ok := ins.(Client); ok {
				out := createStorableClient(clnt)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Client instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if clnt, ok := ins.(Client); ok {
				return []string{
					retrieveAllClientKeyname(),
					retrieveClientByPaidToWalletKeyname(clnt.PaidTo()),
					retrieveClientByHostKeyname(clnt.Host()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Client instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if clnt, ok := ins.(Client); ok {
				// crate metadata and representation:
				metaData := createMetaData()
				hostRepresentation := host.SDKFunc.CreateRepresentation()
				hostMetaData := hostRepresentation.MetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				entityService := entity.SDKFunc.CreateService(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure no other client has the same host:
				retByHst, retByHstErr := repository.RetrieveByHost(clnt.Host())
				if retByHstErr == nil {
					if bytes.Compare(clnt.ID().Bytes(), retByHst.ID().Bytes()) != 0 {
						str := fmt.Sprintf("the given client (ID: %s) contains a host (ID: %s) that is already binded to another client (ID: %s)", clnt.ID().String(), clnt.Host().ID().String(), retByHst.ID().String())
						return errors.New(str)
					}
				}

				// make sure no other client has the same paidTo wallet:
				retByWal, retByWalErr := repository.RetrieveByPaidToWallet(clnt.PaidTo())
				if retByWalErr == nil {
					if bytes.Compare(clnt.ID().Bytes(), retByWal.ID().Bytes()) != 0 {
						str := fmt.Sprintf("the given client (ID: %s) contains a paidTo wallet (ID: %s) that is already binded to another client (ID: %s)", clnt.ID().String(), clnt.PaidTo().ID().String(), retByHst.ID().String())
						return errors.New(str)
					}
				}

				// if the host does not exists, save it:
				_, retHstErr := entityRepository.RetrieveByID(hostMetaData, clnt.Host().ID())
				if retHstErr != nil {
					saveHstErr := entityService.Save(clnt.Host(), hostRepresentation)
					if saveHstErr != nil {
						return saveHstErr
					}
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Client instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
