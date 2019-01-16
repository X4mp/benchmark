package spot

import (
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllSpotKeyname() string {
	return "spots"
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Spot",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableSpot); ok {
				return createSpotFromStorable(storable)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(storableSpot)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createSpotFromStorable(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Spot instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if spt, ok := ins.(Spot); ok {
				out := createStorableSpot(spt)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Spot instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if storable, ok := ins.(*storableSpot); ok {
				return createSpotFromStorable(storable)
			}

			return nil, errors.New("the given storable instance cannot be converted to a Spot instance")
		},
		EmptyNormalized: new(storableSpot),
		EmptyStorable:   new(storableSpot),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if spt, ok := ins.(Spot); ok {
				out := createStorableSpot(spt)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Spot instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if _, ok := ins.(Spot); ok {
				return []string{
					retrieveAllSpotKeyname(),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Spot instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if spt, ok := ins.(Spot); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)

				// make sure there is no other spot with the same ID:
				_, retSptErr := entityRepository.RetrieveByID(metaData, spt.ID())
				if retSptErr == nil {
					str := fmt.Sprintf("there is already a Spot instance that contains this ID: %s", spt.ID().String())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Spot instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
