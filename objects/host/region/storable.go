package region

type storableRegion struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
}

func createStorableRegion(ins Region) *storableRegion {
	out := storableRegion{
		ID:        ins.ID().String(),
		Name:      ins.Name(),
		CountryID: ins.Country().ID().String(),
	}

	return &out
}
