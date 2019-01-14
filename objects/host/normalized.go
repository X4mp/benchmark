package host

import (
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
)

type normalizedHost struct {
	ID           string                  `json:"id"`
	IP           string                  `json:"ip"`
	Hostname     string                  `json:"hostname"`
	Longitude    int                     `json:"longitude"`
	Latitude     int                     `json:"latitude"`
	Country      country.Normalized      `json:"country"`
	Region       region.Normalized       `json:"region"`
	City         city.Normalized         `json:"city"`
	Organization organization.Normalized `json:"organization"`
}

func createNormalizedHost(ins Host) (*normalizedHost, error) {
	var nCountry country.Normalized
	if ins.HasCountry() {
		count, countErr := country.SDKFunc.CreateMetaData().Normalize()(ins.Country())
		if countErr != nil {
			return nil, countErr
		}

		nCountry = count
	}

	var nRegion region.Normalized
	if ins.HasRegion() {
		reg, regErr := region.SDKFunc.CreateMetaData().Normalize()(ins.Region())
		if regErr != nil {
			return nil, regErr
		}

		nRegion = reg
	}

	var nCity city.Normalized
	if ins.HasCity() {
		cit, citErr := city.SDKFunc.CreateMetaData().Normalize()(ins.City())
		if citErr != nil {
			return nil, citErr
		}

		nCity = cit
	}

	var nOrg organization.Normalized
	if ins.HasOrganization() {
		org, orgErr := organization.SDKFunc.CreateMetaData().Normalize()(ins.Organization())
		if orgErr != nil {
			return nil, orgErr
		}

		nOrg = org
	}

	out := normalizedHost{
		ID:           ins.ID().String(),
		IP:           ins.IP().String(),
		Hostname:     ins.Hostname(),
		Longitude:    ins.Longitude(),
		Latitude:     ins.Latitude(),
		Country:      nCountry,
		Region:       nRegion,
		City:         nCity,
		Organization: nOrg,
	}

	return &out, nil
}
