package region

type storableRegion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createStorableRegion(ins Region) *storableRegion {
	out := storableRegion{
		ID:   ins.ID().String(),
		Name: ins.Name(),
	}

	return &out
}
