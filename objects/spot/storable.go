package spot

type storableSpot struct {
	ID        string `json:"id"`
	Longitude int    `json:"longitude"`
	Latitude  int    `json:"latitude"`
	Radius    int    `json:"radius"`
}

func createStorableSpot(ins Spot) *storableSpot {
	out := storableSpot{
		ID:        ins.ID().String(),
		Longitude: ins.Longitude(),
		Latitude:  ins.Latitude(),
		Radius:    ins.Radius(),
	}

	return &out
}
