package report

type storableReport struct {
	ID        string `json:"id"`
	RequestID string `json:"request_id"`
}

func createStorableReport(ins Report) *storableReport {
	out := storableReport{
		ID:        ins.ID().String(),
		RequestID: ins.Request().ID().String(),
	}

	return &out
}
