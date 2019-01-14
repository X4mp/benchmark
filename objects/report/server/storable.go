package server

type storableReport struct {
	ID            string `json:"id"`
	ReportID      string `json:"report_id"`
	ServerID      string `json:"server_id"`
	EncryptedData string `json:"encrypted_data"`
}

func createStorableReport(ins Report) *storableReport {
	out := storableReport{
		ID:            ins.ID().String(),
		ReportID:      ins.Report().ID().String(),
		ServerID:      ins.Server().ID().String(),
		EncryptedData: ins.EncryptedData(),
	}

	return &out
}
