package city

import (
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/region"
)

type normalizedCity struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Region  region.Normalized  `json:"region"`
	Country country.Normalized `json:"country"`
}

func createNormalizedCity(ins City) (*normalizedCity, error) {
	var nReg region.Normalized
	if ins.HasRegion() {
		reg, regErr := region.SDKFunc.CreateMetaData().Normalize()(ins.Region())
		if regErr != nil {
			return nil, regErr
		}

		nReg = reg
	}

	var nCountry country.Normalized
	if ins.HasCountry() {
		count, countErr := country.SDKFunc.CreateMetaData().Normalize()(ins.Country())
		if countErr != nil {
			return nil, countErr
		}

		nCountry = count
	}

	out := normalizedCity{
		ID:      ins.ID().String(),
		Name:    ins.Name(),
		Region:  nReg,
		Country: nCountry,
	}

	return &out, nil

}
