package information

import (
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

type normalizedInformation struct {
	ID                         string            `json:"id"`
	NetworkWallet              wallet.Normalized `json:"network_wallet"`
	PricePerReportPurchase     int               `json:"price_per_report_purchase"`
	RewardPerReport            int               `json:"reward_per_report"`
	MaxSpeedDifferentForNoise  int               `json:"max_speed_different_for_noise"`
	DifferencePercentForStrike int               `json:"difference_percent_for_strike"`
	MaxStrikes                 int               `json:"max_strikes"`
}

func createNormalizedInformation(ins Information) (*normalizedInformation, error) {
	wal, walErr := wallet.SDKFunc.CreateMetaData().Normalize()(ins.NetworkWallet())
	if walErr != nil {
		return nil, walErr
	}

	out := normalizedInformation{
		ID:                         ins.ID().String(),
		NetworkWallet:              wal,
		PricePerReportPurchase:     ins.PricePerReportPurchase(),
		RewardPerReport:            ins.RewardPerReport(),
		MaxSpeedDifferentForNoise:  ins.MaxSpeedDifferentForNoise(),
		DifferencePercentForStrike: ins.DifferencePercentForStrike(),
		MaxStrikes:                 ins.MaxStrikes(),
	}

	return &out, nil
}
