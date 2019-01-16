package information

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllInformationKeyname() string {
	return "benchmarkinformations"
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "BenchmarkInformation",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableInformation); ok {
				return createInformationFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedInformation)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createInformationFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Information instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if inf, ok := ins.(Information); ok {
				return createNormalizedInformation(inf)
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Information instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedInformation); ok {
				return createInformationFromNormalized(normalized)
			}

			return nil, errors.New("the given storable instance cannot be converted to a Information instance")
		},
		EmptyNormalized: new(normalizedInformation),
		EmptyStorable:   new(storableInformation),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if inf, ok := ins.(Information); ok {
				out := createStorableInformation(inf)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Information instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if _, ok := ins.(Information); ok {
				return []string{
					retrieveAllInformationKeyname(),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Information instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if inf, ok := ins.(Information); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				repository := createRepository(metaData, entityRepository)

				// if there is already an information, make sure the ID matches:
				retInf, retInfErr := repository.Retrieve()
				if retInfErr == nil {
					if bytes.Compare(inf.ID().Bytes(), retInf.ID().Bytes()) != 0 {
						str := fmt.Sprintf("the given information instance (ID: %s) has a different ID than the one currently stored (ID: %s)", inf.ID().String(), retInf.ID().String())
						return errors.New(str)
					}
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Information instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
