package request

import (
	"errors"
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllRequestKeyname() string {
	return "benchmarkrequests"
}

func retrieveRequestByClientKeyname(cl client.Client) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_client_id:%s", base, cl.ID().String())
}

func retrieveRequestBySpotKeyname(spt spot.Spot) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_spot_id:%s", base, spt.ID().String())
}

func retrieveRequestByCityKeyname(cit city.City) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_city_id:%s", base, cit.ID().String())
}

func retrieveRequestByRegionKeyname(reg region.Region) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_region_id:%s", base, reg.ID().String())
}

func retrieveRequestByCountryKeyname(count country.Country) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_country_id:%s", base, count.ID().String())
}

func retrieveRequestByOrganizationKeyname(org organization.Organization) string {
	base := retrieveAllRequestKeyname()
	return fmt.Sprintf("%s:by_organization_id:%s", base, org.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "BenchmarkRequest",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableRequest); ok {
				return createRequestFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedRequest)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createRequestFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Request instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if req, ok := ins.(Request); ok {
				out, outErr := createNormalizedRequest(req)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Request instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedRequest); ok {
				return createRequestFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Request instance")
		},
		EmptyStorable:   new(storableRequest),
		EmptyNormalized: new(normalizedRequest),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if req, ok := ins.(Request); ok {
				out := createStorableRequest(req)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Request instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if req, ok := ins.(Request); ok {
				keynames := []string{
					retrieveAllRequestKeyname(),
					retrieveRequestByClientKeyname(req.Client()),
				}

				if req.HasSpots() {
					spots := req.Spots()
					for _, oneSpot := range spots {
						keynames = append(keynames, retrieveRequestBySpotKeyname(oneSpot))
					}
				}

				if req.HasCities() {
					cits := req.Cities()
					for _, oneCity := range cits {
						keynames = append(keynames, retrieveRequestByCityKeyname(oneCity))
					}
				}

				if req.HasRegions() {
					regs := req.Regions()
					for _, oneRegion := range regs {
						keynames = append(keynames, retrieveRequestByRegionKeyname(oneRegion))
					}
				}

				if req.HasCountries() {
					counts := req.Countries()
					for _, oneCountry := range counts {
						keynames = append(keynames, retrieveRequestByCountryKeyname(oneCountry))
					}
				}

				if req.HasOrganizations() {
					orgs := req.Organizations()
					for _, oneOrg := range orgs {
						keynames = append(keynames, retrieveRequestByOrganizationKeyname(oneOrg))
					}
				}

				return keynames, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Request instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if req, ok := ins.(Request); ok {
				// create metadata and representation:
				clientMetaData := client.SDKFunc.CreateMetaData()
				spotMetaData := spot.SDKFunc.CreateMetaData()
				cityMetaData := city.SDKFunc.CreateMetaData()
				regionMetaData := region.SDKFunc.CreateMetaData()
				countryMetaData := country.SDKFunc.CreateMetaData()
				organizationMetaData := organization.SDKFunc.CreateMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)

				// make sure the client exists:
				_, retClientErr := entityRepository.RetrieveByID(clientMetaData, req.Client().ID())
				if retClientErr != nil {
					str := fmt.Sprintf("the Request (ID: %s) contains a Client (ID: %s) that does not exists", req.ID().String(), req.Client().ID().String())
					return errors.New(str)
				}

				if req.HasSpots() {
					spts := req.Spots()
					for _, oneSpot := range spts {
						_, retSpotErr := entityRepository.RetrieveByID(spotMetaData, oneSpot.ID())
						if retSpotErr != nil {
							str := fmt.Sprintf("the Request (ID: %s) contains a Client (ID: %s) that does not exists", req.ID().String(), oneSpot.ID().String())
							return errors.New(str)
						}
					}
				}

				if req.HasCities() {
					cits := req.Cities()
					for _, oneCity := range cits {
						_, retCityErr := entityRepository.RetrieveByID(cityMetaData, oneCity.ID())
						if retCityErr != nil {
							str := fmt.Sprintf("the Request (ID: %s) contains a City (ID: %s) that does not exists", req.ID().String(), oneCity.ID().String())
							return errors.New(str)
						}
					}
				}

				if req.HasRegions() {
					regs := req.Regions()
					for _, oneRegion := range regs {
						_, retRegErr := entityRepository.RetrieveByID(regionMetaData, oneRegion.ID())
						if retRegErr != nil {
							str := fmt.Sprintf("the Request (ID: %s) contains a Region (ID: %s) that does not exists", req.ID().String(), oneRegion.ID().String())
							return errors.New(str)
						}
					}
				}

				if req.HasCountries() {
					counts := req.Countries()
					for _, oneCountry := range counts {
						_, retCountryErr := entityRepository.RetrieveByID(countryMetaData, oneCountry.ID())
						if retCountryErr != nil {
							str := fmt.Sprintf("the Request (ID: %s) contains a Country (ID: %s) that does not exists", req.ID().String(), oneCountry.ID().String())
							return errors.New(str)
						}
					}
				}

				if req.HasOrganizations() {
					orgs := req.Organizations()
					for _, oneOrg := range orgs {
						_, retOrgErr := entityRepository.RetrieveByID(organizationMetaData, oneOrg.ID())
						if retOrgErr != nil {
							str := fmt.Sprintf("the Request (ID: %s) contains an Organization (ID: %s) that does not exists", req.ID().String(), oneOrg.ID().String())
							return errors.New(str)
						}
					}
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Request instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
