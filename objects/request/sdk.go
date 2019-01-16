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

// Request represents a client request
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
