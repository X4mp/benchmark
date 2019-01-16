package request

import (
	"time"

	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
)

type normalizedRequest struct {
	ID            string                    `json:"id"`
	Client        client.Normalized         `json:"client"`
	Interval      *time.Duration            `json:"interval"`
	Spots         []spot.Normalized         `json:"spots"`
	Cities        []city.Normalized         `json:"cities"`
	Regions       []region.Normalized       `json:"regions"`
	Countries     []country.Normalized      `json:"countries"`
	Organizations []organization.Normalized `json:"organizations"`
}

func createNormalizedRequest(ins Request) (*normalizedRequest, error) {
	cl, clErr := client.SDKFunc.CreateMetaData().Normalize()(ins.Client())
	if clErr != nil {
		return nil, clErr
	}

	nSpots := []spot.Normalized{}
	if ins.HasSpots() {
		spots := ins.Spots()
		for _, oneSpot := range spots {
			nOneSpot, nOneSpotErr := spot.SDKFunc.CreateMetaData().Normalize()(oneSpot)
			if nOneSpotErr != nil {
				return nil, nOneSpotErr
			}

			nSpots = append(nSpots, nOneSpot)
		}
	}

	nCities := []city.Normalized{}
	if ins.HasCities() {
		cities := ins.Cities()
		for _, oneCity := range cities {
			nOneCity, nOneCityErr := city.SDKFunc.CreateMetaData().Normalize()(oneCity)
			if nOneCityErr != nil {
				return nil, nOneCityErr
			}

			nCities = append(nCities, nOneCity)
		}
	}

	nRegions := []region.Normalized{}
	if ins.HasRegions() {
		regions := ins.Regions()
		for _, oneRegion := range regions {
			nOneRegion, nOneRegionErr := region.SDKFunc.CreateMetaData().Normalize()(oneRegion)
			if nOneRegionErr != nil {
				return nil, nOneRegionErr
			}

			nRegions = append(nRegions, nOneRegion)
		}
	}

	nCountries := []country.Normalized{}
	if ins.HasCountries() {
		countries := ins.Countries()
		for _, oneCountry := range countries {
			nOneCountry, nOneCountryErr := country.SDKFunc.CreateMetaData().Normalize()(oneCountry)
			if nOneCountryErr != nil {
				return nil, nOneCountryErr
			}

			nCountries = append(nCountries, nOneCountry)
		}
	}

	nOrgs := []organization.Normalized{}
	if ins.HasOrganizations() {
		orgs := ins.Organizations()
		for _, oneOrg := range orgs {
			nOneOrg, nOneOrgErr := organization.SDKFunc.CreateMetaData().Normalize()(oneOrg)
			if nOneOrgErr != nil {
				return nil, nOneOrgErr
			}

			nOrgs = append(nOrgs, nOneOrg)
		}
	}

	out := normalizedRequest{
		ID:            ins.ID().String(),
		Client:        cl,
		Interval:      ins.Interval(),
		Spots:         nSpots,
		Cities:        nCities,
		Regions:       nRegions,
		Countries:     nCountries,
		Organizations: nOrgs,
	}

	return &out, nil
}
