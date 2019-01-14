package report

type storableReport struct {
	ID       string `json:"id"`
	ClientID string `json:"client_id"`
}

func createStorableReport(ins Report) *storableReport {
	out := storableReport{
		ID:       ins.ID().String(),
		ClientID: ins.Client().ID().String(),
	}

	return &out
}
