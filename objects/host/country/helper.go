package country

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllCountryKeyname() string {
	return "counties"
}

func retrieveCountryByNameKeyname(name string) string {
	h := sha256.New()
	h.Write([]byte(name))
	base := retrieveAllCountryKeyname()
	return fmt.Sprintf("%s:by_hashed_name:%s", base, hex.EncodeToString(h.Sum(nil)))
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Country",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableCountry); ok {
				return createCountryFromStorable(storable)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(storableCountry)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createCountryFromStorable(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Country instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if count, ok := ins.(Country); ok {
				out := createStorableCountry(count)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Country instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if storable, ok := ins.(*storableCountry); ok {
				return createCountryFromStorable(storable)
			}

			return nil, errors.New("the given storable instance cannot be converted to a Country instance")
		},
		EmptyNormalized: new(storableCountry),
		EmptyStorable:   new(storableCountry),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if count, ok := ins.(Country); ok {
				out := createStorableCountry(count)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Country instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if count, ok := ins.(Country); ok {
				return []string{
					retrieveAllCountryKeyname(),
					retrieveCountryByNameKeyname(count.Name()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if count, ok := ins.(Country); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure the name is unique:
				_, retCountryErr := repository.RetrieveByName(count.Name())
				if retCountryErr == nil {
					str := fmt.Sprintf("the country (Name: %s) already exists", count.Name())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Country instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
