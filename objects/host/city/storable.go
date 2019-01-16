package city

type storableCity struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	RegionID  string `json:"region_id"`
	CountryID string `json:"country_id"`
}

func createStorableCity(ins City) *storableCity {
	regionID := ""
	if ins.HasRegion() {
		regionID = ins.Region().ID().String()
	}

	countryID := ""
	if ins.HasCountry() {
		countryID = ins.Country().ID().String()
	}

	out := storableCity{
		ID:        ins.ID().String(),
		Name:      ins.Name(),
		RegionID:  regionID,
		CountryID: countryID,
	}

	return &out
}
