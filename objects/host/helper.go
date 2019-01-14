package host

import (
	"errors"
	"fmt"
	"net"

	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllHostKeyname() string {
	return "hosts"
}

func retrieveHostByIPKeyname(ip net.IP) string {
	base := retrieveAllHostKeyname()
	return fmt.Sprintf("%s:by_ip:%s", base, ip.String())
}

func retrieveHostByCountryKeyname(count country.Country) string {
	base := retrieveAllHostKeyname()
	return fmt.Sprintf("%s:by_country_id:%s", base, count.ID().String())
}

func retrieveHostByRegionKeyname(reg region.Region) string {
	base := retrieveAllHostKeyname()
	return fmt.Sprintf("%s:by_region_id:%s", base, reg.ID().String())
}

func retrieveHostByCityKeyname(cit city.City) string {
	base := retrieveAllHostKeyname()
	return fmt.Sprintf("%s:by_city_id:%s", base, cit.ID().String())
}

func retrieveHostByOrganizationKeyname(org organization.Organization) string {
	base := retrieveAllHostKeyname()
	return fmt.Sprintf("%s:by_organization_id:%s", base, org.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Host",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableHost); ok {
				return createHostFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedHost)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createHostFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Host instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if hst, ok := ins.(Host); ok {
				out, outErr := createNormalizedHost(hst)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Host instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedHost); ok {
				return createHostFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Host instance")
		},
		EmptyStorable:   new(storableHost),
		EmptyNormalized: new(normalizedHost),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if hst, ok := ins.(Host); ok {
				out := createStorableHost(hst)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Host instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if hst, ok := ins.(Host); ok {
				return []string{
					retrieveAllHostKeyname(),
					retrieveHostByIPKeyname(hst.IP()),
					retrieveHostByCountryKeyname(hst.Country()),
					retrieveHostByRegionKeyname(hst.Region()),
					retrieveHostByCityKeyname(hst.City()),
					retrieveHostByOrganizationKeyname(hst.Organization()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Host instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if hst, ok := ins.(Host); ok {
				// crate metadata and representation:
				metaData := createMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)

				// make sure the host doesnt already exists:
				_, retHostErr := entityRepository.RetrieveByID(metaData, hst.ID())
				if retHostErr == nil {
					str := fmt.Sprintf("the host (ID: %s) already exists", hst.ID().String())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Host instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
