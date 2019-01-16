package region

import (
	"github.com/xmnnetwork/benchmark/objects/host/country"
)

type normalizedRegion struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Country country.Normalized `json:"country"`
}

func createNormalizedRegion(ins Region) (*normalizedRegion, error) {
	count, countErr := country.SDKFunc.CreateMetaData().Normalize()(ins.Country())
	if countErr != nil {
		return nil, countErr
	}

	out := normalizedRegion{
		ID:      ins.ID().String(),
		Name:    ins.Name(),
		Country: count,
	}

	return &out, nil
}
