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

func createHostWithCity(
	id *uuid.UUID,
	ipAddress net.IP,
	hostName string,
	long int,
	lat int,
	org organization.Organization,
	city city.City,
) (Host, error) {
	out := host{
		UUID:      id,
		IPAddress: ipAddress,
		HstName:   hostName,
		Long:      long,
		Lat:       lat,
		Org:       org,
		Cit:       city,
		Reg:       nil,
		Count:     nil,
	}

	return &out, nil
}

func createHostWithRegion(
	id *uuid.UUID,
	ipAddress net.IP,
	hostName string,
	long int,
	lat int,
	org organization.Organization,
	region region.Region,
) (Host, error) {
	out := host{
		UUID:      id,
		IPAddress: ipAddress,
		HstName:   hostName,
		Long:      long,
		Lat:       lat,
		Org:       org,
		Cit:       nil,
		Reg:       region,
		Count:     nil,
	}

	return &out, nil
}

func createHostWithCountry(
	id *uuid.UUID,
	ipAddress net.IP,
	hostName string,
	long int,
	lat int,
	org organization.Organization,
	country country.Country,
) (Host, error) {
	out := host{
		UUID:      id,
		IPAddress: ipAddress,
		HstName:   hostName,
		Long:      long,
		Lat:       lat,
		Org:       org,
		Cit:       nil,
		Reg:       nil,
		Count:     country,
	}

	return &out, nil
}

func createHostFromNormalized(normalized *normalizedHost) (Host, error) {
	ipAddress := net.ParseIP(normalized.IP)
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	var org organization.Organization
	if normalized.Organization != nil {
		orgIns, orgInsErr := organization.SDKFunc.CreateMetaData().Denormalize()(normalized.Organization)
		if orgInsErr != nil {
			return nil, orgInsErr
		}

		if casted, ok := orgIns.(organization.Organization); ok {
			org = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", orgIns.ID().String())
		return nil, errors.New(str)
	}

	if normalized.City != nil {
		cityIns, cityInsErr := city.SDKFunc.CreateMetaData().Denormalize()(normalized.City)
		if cityInsErr != nil {
			return nil, cityInsErr
		}

		if cit, ok := cityIns.(city.City); ok {
			return createHostWithCity(&id, ipAddress, normalized.Hostname, normalized.Latitude, normalized.Longitude, org, cit)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", cityIns.ID().String())
		return nil, errors.New(str)
	}

	if normalized.Region != nil {
		regIns, regInsErr := region.SDKFunc.CreateMetaData().Denormalize()(normalized.Region)
		if regInsErr != nil {
			return nil, regInsErr
		}

		if reg, ok := regIns.(region.Region); ok {
			return createHostWithRegion(&id, ipAddress, normalized.Hostname, normalized.Latitude, normalized.Longitude, org, reg)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regIns.ID().String())
		return nil, errors.New(str)
	}

	counIns, counInsErr := region.SDKFunc.CreateMetaData().Denormalize()(normalized.Region)
	if counInsErr != nil {
		return nil, counInsErr
	}

	if count, ok := counIns.(country.Country); ok {
		return createHostWithCountry(&id, ipAddress, normalized.Hostname, normalized.Latitude, normalized.Longitude, org, count)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", counIns.ID().String())
	return nil, errors.New(str)
}

func createHostFromStorable(storable *storableHost, rep entity.Repository) (Host, error) {
	ipAddress := net.ParseIP(storable.IP)
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
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
			org = casted
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Organization instance", orgIns.ID().String())
		return nil, errors.New(str)
	}

	if storable.CityID != "" {
		cityID, cityIDErr := uuid.FromString(storable.CityID)
		if cityIDErr != nil {
			return nil, cityIDErr
		}

		cityIns, cityInsErr := rep.RetrieveByID(city.SDKFunc.CreateMetaData(), &cityID)
		if cityInsErr != nil {
			return nil, cityInsErr
		}

		if cit, ok := cityIns.(city.City); ok {
			return createHostWithCity(&id, ipAddress, storable.Hostname, storable.Latitude, storable.Longitude, org, cit)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid City instance", cityIns.ID().String())
		return nil, errors.New(str)
	}

	if storable.RegionID != "" {
		regionID, regionIDErr := uuid.FromString(storable.RegionID)
		if regionIDErr != nil {
			return nil, regionIDErr
		}

		regIns, regInsErr := rep.RetrieveByID(region.SDKFunc.CreateMetaData(), &regionID)
		if regInsErr != nil {
			return nil, regInsErr
		}

		if reg, ok := regIns.(region.Region); ok {
			return createHostWithRegion(&id, ipAddress, storable.Hostname, storable.Latitude, storable.Longitude, org, reg)
		}

		str := fmt.Sprintf("the entity (ID: %s) is not a valid Region instance", regIns.ID().String())
		return nil, errors.New(str)
	}

	countID, countIDErr := uuid.FromString(storable.CountryID)
	if countIDErr != nil {
		return nil, countIDErr
	}

	counIns, counInsErr := rep.RetrieveByID(country.SDKFunc.CreateMetaData(), &countID)
	if counInsErr != nil {
		return nil, counInsErr
	}

	if count, ok := counIns.(country.Country); ok {
		return createHostWithCountry(&id, ipAddress, storable.Hostname, storable.Latitude, storable.Longitude, org, count)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Country instance", counIns.ID().String())
	return nil, errors.New(str)
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
