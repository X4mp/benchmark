package city

type storableCity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createStorableCity(ins City) *storableCity {
	out := storableCity{
		ID:   ins.ID().String(),
		Name: ins.Name(),
	}

	return &out
}
