package customer

type storableCustomer struct {
	ID         string `json:"id"`
	TransferID string `json:"transfer_id"`
	ReportID   string `json:"report_id"`
}

func createStorableCustomer(ins Customer) *storableCustomer {
	out := storableCustomer{
		ID:         ins.ID().String(),
		TransferID: ins.Transfer().ID().String(),
		ReportID:   ins.Report().ID().String(),
	}

	return &out
}
