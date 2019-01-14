package country

type storableCountry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createStorableCountry(ins Country) *storableCountry {
	out := storableCountry{
		ID:   ins.ID().String(),
		Name: ins.Name(),
	}

	return &out
}
