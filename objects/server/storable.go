package server

type storableServer struct {
	ID       string `json:"id"`
	PledgeID string `json:"pledge_id"`
	HostID   string `json:"host_id"`
}

func createStorableServer(ins Server) *storableServer {
	out := storableServer{
		ID:       ins.ID().String(),
		PledgeID: ins.Pledge().ID().String(),
		HostID:   ins.Host().ID().String(),
	}

	return &out
}
