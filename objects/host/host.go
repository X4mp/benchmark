package host

import (
	"errors"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

type host struct {
	UUID      *uuid.UUID                `json:"id"`
	IPAddress net.IP                    `json:"ip"`
	HstName   string                    `json:"hostname"`
	Cit       city.City                 `json:"city"`
	Reg       region.Region             `json:"region"`
	Count     country.Country           `json:"country"`
	Long      int                       `json:"longitude"`
	Lat       int                       `json:"latitude"`
	Org       organization.Organization `json:"organization"`
}

func createHost(
	id *uuid.UUID,
	ipAddress net.IP,
	hostName string,
	city city.City,
	region region.Region,
	country country.Country,
	long int,
	lat int,
	org organization.Organization,
) (Host, error) {
	out := host{
		UUID:      id,
		IPAddress: ipAddress,
		HstName:   hostName,
		Cit:       city,
		Reg:       region,
		Count:     country,
		Long:      long,
		Lat:       lat,
		Org:       org,
	}

	return &out, nil
}

func createHostFromNormalized(normalized *normalizedHost) (Host, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	var cit city.City
	if normalized.City != nil {
		cityIns, cityInsErr := city.SDKFunc.CreateMetaData().Denormalize()(normalized.City)
		if cityInsErr != nil {
			return nil, cityInsErr
		}

		if casted, ok := cityIns.(city.City); ok {
			cit = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", cityIns.ID().String())
		return nil, errors.New(str)
	}

	var reg region.Region
	if normalized.Region != nil {
		regIns, regInsErr := region.SDKFunc.CreateMetaData().Denormalize()(normalized.Region)
		if regInsErr != nil {
			return nil, regInsErr
		}

		if casted, ok := regIns.(region.Region); ok {
			reg = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regIns.ID().String())
		return nil, errors.New(str)
	}

	var count country.Country
	if normalized.Country != nil {
		counIns, counInsErr := region.SDKFunc.CreateMetaData().Denormalize()(normalized.Region)
		if counInsErr != nil {
			return nil, counInsErr
		}

		if casted, ok := counIns.(country.Country); ok {
			count = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", counIns.ID().String())
		return nil, errors.New(str)
	}

	var org organization.Organization
	if normalized.Organization != nil {
		orgIns, orgInsErr := organization.SDKFunc.CreateMetaData().Denormalize()(normalized.Organization)
		if orgInsErr != nil {
			return nil, orgInsErr
		}

		if casted, ok := orgIns.(organization.Organization); ok {
			count = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", orgIns.ID().String())
		return nil, errors.New(str)
	}

	ipAddress := net.ParseIP(normalized.IP)
	return createHost(&id, ipAddress, normalized.Hostname, cit, reg, count, normalized.Latitude, normalized.Longitude, org)
}

func createHostFromStorable(storable *storableHost, rep entity.Repository) (Host, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	var cit city.City
	if storable.CityID != "" {
		cityID, cityIDErr := uuid.FromString(storable.CityID)
		if cityIDErr != nil {
			return nil, cityIDErr
		}

		cityIns, cityInsErr := rep.RetrieveByID(city.SDKFunc.CreateMetaData(), &cityID)
		if cityInsErr != nil {
			return nil, cityInsErr
		}

		if casted, ok := cityIns.(city.City); ok {
			cit = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", cityIns.ID().String())
		return nil, errors.New(str)
	}

	var reg region.Region
	if storable.RegionID != "" {
		regionID, regionIDErr := uuid.FromString(storable.RegionID)
		if regionIDErr != nil {
			return nil, regionIDErr
		}

		regIns, regInsErr := rep.RetrieveByID(region.SDKFunc.CreateMetaData(), &regionID)
		if regInsErr != nil {
			return nil, regInsErr
		}

		if casted, ok := regIns.(region.Region); ok {
			reg = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regIns.ID().String())
		return nil, errors.New(str)
	}

	var count country.Country
	if storable.CountryID != "" {
		countID, countIDErr := uuid.FromString(storable.CountryID)
		if countIDErr != nil {
			return nil, countIDErr
		}

		counIns, counInsErr := rep.RetrieveByID(country.SDKFunc.CreateMetaData(), &countID)
		if counInsErr != nil {
			return nil, counInsErr
		}

		if casted, ok := counIns.(country.Country); ok {
			count = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", counIns.ID().String())
		return nil, errors.New(str)
	}

	var org organization.Organization
	if storable.OrganizationID != "" {
		orgID, orgIDErr := uuid.FromString(storable.OrganizationID)
		if orgIDErr != nil {
			return nil, orgIDErr
		}

		orgIns, orgInsErr := rep.RetrieveByID(organization.SDKFunc.CreateMetaData(), &orgID)
		if orgInsErr != nil {
			return nil, orgInsErr
		}

		if casted, ok := orgIns.(organization.Organization); ok {
			count = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", orgIns.ID().String())
		return nil, errors.New(str)
	}

	ipAddress := net.ParseIP(storable.IP)
	return createHost(&id, ipAddress, storable.Hostname, cit, reg, count, storable.Latitude, storable.Longitude, org)
}

// ID returns the ID
func (obj *host) ID() *uuid.UUID {
	return obj.UUID
}

// IP returns the IP
func (obj *host) IP() net.IP {
	return obj.IPAddress
}

// Hostname returns the hostname
func (obj *host) Hostname() string {
	return obj.HstName
}

// Longitude returns the longitude
func (obj *host) Longitude() int {
	return obj.Long
}

// Latitude returns the latitude
func (obj *host) Latitude() int {
	return obj.Lat
}

// HasCity returns true if there is a city
func (obj *host) HasCity() bool {
	return obj.Cit != nil
}

// City returns the city
func (obj *host) City() city.City {
	return obj.Cit
}

// HasRegion returns true if there is a region
func (obj *host) HasRegion() bool {
	return obj.Reg != nil
}

// Region returns the region
func (obj *host) Region() region.Region {
	return obj.Reg
}

// HasCountry returns true if there is a country
func (obj *host) HasCountry() bool {
	return obj.Count != nil
}

// Country returns the country
func (obj *host) Country() country.Country {
	return obj.Count
}

// HasOrganization returns true if there is an organization
func (obj *host) HasOrganization() bool {
	return obj.Org != nil
}

// Organization returns the organization
func (obj *host) Organization() organization.Organization {
	return obj.Org
}
