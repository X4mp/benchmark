package city

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllCityKeyname() string {
	return "cities"
}

func retrieveCityByNameKeyname(name string) string {
	h := sha256.New()
	h.Write([]byte(name))
	base := retrieveAllCityKeyname()
	return fmt.Sprintf("%s:by_hashed_name:%s", base, hex.EncodeToString(h.Sum(nil)))
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "City",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableCity); ok {
				return createCityFromStorable(storable)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(storableCity)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createCityFromStorable(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a City instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if cit, ok := ins.(City); ok {
				out := createStorableCity(cit)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid City instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if storable, ok := ins.(*storableCity); ok {
				return createCityFromStorable(storable)
			}

			return nil, errors.New("the given storable instance cannot be converted to a City instance")
		},
		EmptyNormalized: new(storableCity),
		EmptyStorable:   new(storableCity),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if cit, ok := ins.(City); ok {
				out := createStorableCity(cit)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid City instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if cit, ok := ins.(City); ok {
				return []string{
					retrieveAllCityKeyname(),
					retrieveCityByNameKeyname(cit.Name()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if cit, ok := ins.(City); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure the name is unique:
				_, retCityErr := repository.RetrieveByName(cit.Name())
				if retCityErr == nil {
					str := fmt.Sprintf("the city (Name: %s) already exists", cit.Name())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid City instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
