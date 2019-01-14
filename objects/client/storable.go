package client

type storableClient struct {
	ID             string `json:"id"`
	PaidToWalletID string `json:"paid_to_wallet_id"`
	HostID         string `json:"host_id"`
}

func createStorableClient(ins Client) *storableClient {
	out := storableClient{
		ID:             ins.ID().String(),
		PaidToWalletID: ins.PaidTo().ID().String(),
		HostID:         ins.Host().ID().String(),
	}

	return &out
}
