package information

type storableInformation struct {
	ID                         string `json:"id"`
	PricePerReportPurchase     int    `json:"price_per_report_purchase"`
	RewardPerReport            int    `json:"reward_per_report"`
	MaxSpeedDifferentForNoise  int    `json:"max_speed_different_for_noise"`
	DifferencePercentForStrike int    `json:"difference_percent_for_strike"`
	MaxStrikes                 int    `json:"max_strikes"`
}

func createStorableInformation(ins Information) *storableInformation {
	out := storableInformation{
		ID: ins.ID().String(),
		PricePerReportPurchase:     ins.PricePerReportPurchase(),
		RewardPerReport:            ins.RewardPerReport(),
		MaxSpeedDifferentForNoise:  ins.MaxSpeedDifferentForNoise(),
		DifferencePercentForStrike: ins.DifferencePercentForStrike(),
		MaxStrikes:                 ins.MaxStrikes(),
	}

	return &out
}
