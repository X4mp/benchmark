package server

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllServerKeyname() string {
	return "servers"
}

func retrieveServerByPledgeKeyname(pldge pledge.Pledge) string {
	base := retrieveAllServerKeyname()
	return fmt.Sprintf("%s:by_pledge_id:%s", base, pldge.ID().String())
}

func retrieveServerByHostKeyname(hst host.Host) string {
	base := retrieveAllServerKeyname()
	return fmt.Sprintf("%s:by_host_id:%s", base, hst.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Server",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableServer); ok {
				return createServerFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedServer)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createServerFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Server instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if serv, ok := ins.(Server); ok {
				out, outErr := createNormalizedServer(serv)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Server instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedServer); ok {
				return createServerFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Server instance")
		},
		EmptyStorable:   new(storableServer),
		EmptyNormalized: new(normalizedServer),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if serv, ok := ins.(Server); ok {
				out := createStorableServer(serv)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Server instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if serv, ok := ins.(Server); ok {
				return []string{
					retrieveAllServerKeyname(),
					retrieveServerByPledgeKeyname(serv.Pledge()),
					retrieveServerByHostKeyname(serv.Host()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Server instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if serv, ok := ins.(Server); ok {
				// crate metadata and representation:
				metaData := createMetaData()
				hostRepresentation := host.SDKFunc.CreateRepresentation()
				hostMetaData := hostRepresentation.MetaData()
				pledgeRepresentation := pledge.SDKFunc.CreateRepresentation()
				pledgeMetaData := pledgeRepresentation.MetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				entityService := entity.SDKFunc.CreateService(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure no other client has the same host:
				retByHst, retByHstErr := repository.RetrieveByHost(serv.Host())
				if retByHstErr == nil {
					if bytes.Compare(serv.ID().Bytes(), retByHst.ID().Bytes()) != 0 {
						str := fmt.Sprintf("the given client (ID: %s) contains a host (ID: %s) that is already binded to another client (ID: %s)", serv.ID().String(), serv.Host().ID().String(), retByHst.ID().String())
						return errors.New(str)
					}
				}

				// if the host does not exists, save it:
				_, retHstErr := entityRepository.RetrieveByID(hostMetaData, serv.Host().ID())
				if retHstErr != nil {
					saveHstErr := entityService.Save(serv.Host(), hostRepresentation)
					if saveHstErr != nil {
						return saveHstErr
					}
				}

				// make sure the pledge does not already exists:
				_, retPledgeErr := entityRepository.RetrieveByID(pledgeMetaData, serv.Pledge().ID())
				if retPledgeErr == nil {
					str := fmt.Sprintf("the server (ID: %s) contains a Pledge (ID: %s) that already exists", serv.ID().String(), serv.Pledge().ID().String())
					return errors.New(str)
				}

				// save the pledge:
				savePledgeErr := entityService.Save(serv.Pledge(), pledgeRepresentation)
				if savePledgeErr != nil {
					return savePledgeErr
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Server instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
