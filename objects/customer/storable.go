package customer

type storableCustomer struct {
	ID         string   `json:"id"`
	TransferID string   `json:"transfer_id"`
	ReportIDs  []string `json:"report_ids"`
}

func createStorableCustomer(ins Customer) *storableCustomer {
	reportIDs := []string{}
	reps := ins.Reports()
	for _, oneReport := range reps {
		reportIDs = append(reportIDs, oneReport.ID().String())
	}

	out := storableCustomer{
		ID:         ins.ID().String(),
		TransferID: ins.Transfer().ID().String(),
		ReportIDs:  reportIDs,
	}

	return &out
}
