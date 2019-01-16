package request

import (
	"time"
)

type storableRequest struct {
	ID              string         `json:"id"`
	ClientID        string         `json:"client"`
	Interval        *time.Duration `json:"interval"`
	SpotIDs         []string       `json:"spots"`
	CityIDs         []string       `json:"cities"`
	RegionIDs       []string       `json:"regions"`
	CountryIDs      []string       `json:"countries"`
	OrganizationIDs []string       `json:"organizations"`
}

func createStorableRequest(ins Request) *storableRequest {

	var inter *time.Duration
	if ins.HasInterval() {
		inter = ins.Interval()
	}

	spotIDs := []string{}
	if ins.HasSpots() {
		spots := ins.Spots()
		for _, oneSpot := range spots {
			spotIDs = append(spotIDs, oneSpot.ID().String())
		}
	}

	cityIDs := []string{}
	if ins.HasCities() {
		cities := ins.Cities()
		for _, oneCity := range cities {
			cityIDs = append(cityIDs, oneCity.ID().String())
		}
	}

	regionIDs := []string{}
	if ins.HasRegions() {
		regions := ins.Regions()
		for _, oneRegion := range regions {
			regionIDs = append(regionIDs, oneRegion.ID().String())
		}
	}

	countryIDs := []string{}
	if ins.HasCountries() {
		countries := ins.Countries()
		for _, oneCountry := range countries {
			countryIDs = append(countryIDs, oneCountry.ID().String())
		}
	}

	orgIDs := []string{}
	if ins.HasOrganizations() {
		orgs := ins.Organizations()
		for _, oneOrg := range orgs {
			orgIDs = append(orgIDs, oneOrg.ID().String())
		}
	}

	out := storableRequest{
		ID:              ins.ID().String(),
		ClientID:        ins.Client().ID().String(),
		Interval:        inter,
		SpotIDs:         spotIDs,
		CityIDs:         cityIDs,
		RegionIDs:       regionIDs,
		CountryIDs:      countryIDs,
		OrganizationIDs: orgIDs,
	}

	return &out
}
