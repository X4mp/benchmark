package organization

type storableOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createStorableOrganization(ins Organization) *storableOrganization {
	out := storableOrganization{
		ID:   ins.ID().String(),
		Name: ins.Name(),
	}

	return &out
}
