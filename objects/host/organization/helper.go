package organization

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllOrganizationKeyname() string {
	return "orgies"
}

func retrieveOrganizationByNameKeyname(name string) string {
	h := sha256.New()
	h.Write([]byte(name))
	base := retrieveAllOrganizationKeyname()
	return fmt.Sprintf("%s:by_hashed_name:%s", base, hex.EncodeToString(h.Sum(nil)))
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Organization",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableOrganization); ok {
				return createOrganizationFromStorable(storable)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(storableOrganization)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createOrganizationFromStorable(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Organization instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if org, ok := ins.(Organization); ok {
				out := createStorableOrganization(org)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Organization instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if storable, ok := ins.(*storableOrganization); ok {
				return createOrganizationFromStorable(storable)
			}

			return nil, errors.New("the given storable instance cannot be converted to a Organization instance")
		},
		EmptyNormalized: new(storableOrganization),
		EmptyStorable:   new(storableOrganization),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if org, ok := ins.(Organization); ok {
				out := createStorableOrganization(org)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Organization instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if org, ok := ins.(Organization); ok {
				return []string{
					retrieveAllOrganizationKeyname(),
					retrieveOrganizationByNameKeyname(org.Name()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if org, ok := ins.(Organization); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				repository := createRepository(metaData, entityRepository)

				// make sure the name is unique:
				_, retOrganizationErr := repository.RetrieveByName(org.Name())
				if retOrganizationErr == nil {
					str := fmt.Sprintf("the organization (Name: %s) already exists", org.Name())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Organization instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
