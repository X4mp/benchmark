package region

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllRegionKeyname() string {
	return "regies"
}

func retrieveRegionByNameKeyname(name string) string {
	h := sha256.New()
	h.Write([]byte(name))
	base := retrieveAllRegionKeyname()
	return fmt.Sprintf("%s:by_hashed_name:%s", base, hex.EncodeToString(h.Sum(nil)))
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Region",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableRegion); ok {
				return createRegionFromStorable(storable)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(storableRegion)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createRegionFromStorable(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Region instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if reg, ok := ins.(Region); ok {
				out := createStorableRegion(reg)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Region instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if storable, ok := ins.(*storableRegion); ok {
				return createRegionFromStorable(storable)
			}

			return nil, errors.New("the given storable instance cannot be converted to a Region instance")
		},
		EmptyNormalized: new(storableRegion),
		EmptyStorable:   new(storableRegion),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if reg, ok := ins.(Region); ok {
				out := createStorableRegion(reg)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Region instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if reg, ok := ins.(Region); ok {
				return []string{
					retrieveAllRegionKeyname(),
					retrieveRegionByNameKeyname(reg.Name()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if reg, ok := ins.(Region); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure the name is unique:
				_, retRegionErr := repository.RetrieveByName(reg.Name())
				if retRegionErr == nil {
					str := fmt.Sprintf("the region (Name: %s) already exists", reg.Name())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Region instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
