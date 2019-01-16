package request

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
)

type request struct {
	UUID    *uuid.UUID                  `json:"id"`
	Clnt    client.Client               `json:"client"`
	Int     time.Duration               `json:"interval"`
	Spts    []spot.Spot                 `json:"spots"`
	Cits    []city.City                 `json:"cities"`
	Regs    []region.Region             `json:"regions"`
	Countrs []country.Country           `json:"countries"`
	Orgs    []organization.Organization `json:"organizations"`
}

func createRequest(
	id *uuid.UUID,
	clint client.Client,
	intrvl time.Duration,
	spts []spot.Spot,
	cities []city.City,
	regions []region.Region,
	countries []country.Country,
	orgs []organization.Organization,
) (Request, error) {

}

type Request interface {
	ID() *uuid.UUID
	Client() client.Client
	HasInterval() bool
	Interval() time.Duration
	HasSpot() bool
	Spots() []spot.Spot
	HasCities() bool
	Cities() []city.City
	HasRegions() bool
	Regions() []region.Region
	HasCountries() bool
	Countries() []country.Country
	HasOrganizations() bool
	Organizations() []organization.Organization
}
