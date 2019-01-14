package information

import uuid "github.com/satori/go.uuid"

type information struct {
	UUID                 *uuid.UUID `json:"id"`
	Price                int        `json:"price_per_report_purchase"`
	Reward               int        `json:"reward_per_report"`
	MaxSpeedDiff         int        `json:"max_speed_difference_for_noise"`
	DiffPercentForStrike int        `json:"difference_for_strike"`
	MxStrikes            int        `json:"max_strikes"`
}

func createInformation(id *uuid.UUID, price int, reward int, maxSpeedDiff int, diffPercentForStrike int, maxStrikes int) (Information, error) {
	out := information{
		UUID:                 id,
		Price:                price,
		Reward:               reward,
		MaxSpeedDiff:         maxSpeedDiff,
		DiffPercentForStrike: diffPercentForStrike,
		MxStrikes:            maxStrikes,
	}

	return &out, nil
}

func createInformationFromStorable(storable *storableInformation) (Information, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createInformation(
		&id,
		storable.PricePerReportPurchase,
		storable.RewardPerReport,
		storable.MaxSpeedDifferentForNoise,
		storable.DifferencePercentForStrike,
		storable.MaxStrikes,
	)
}

// ID returns the ID
func (obj *information) ID() *uuid.UUID {
	return obj.UUID
}

// PricePerReportPurchase returns the price per report purchase
func (obj *information) PricePerReportPurchase() int {
	return obj.Price
}

// RewardPerReport returns the reward per report
func (obj *information) RewardPerReport() int {
	return obj.Reward
}

// MaxSpeedDifferentForNoise returns the max speed different for noise
func (obj *information) MaxSpeedDifferentForNoise() int {
	return obj.MaxSpeedDiff
}

// DifferencePercentForStrike returns the difference percent for strike
func (obj *information) DifferencePercentForStrike() int {
	return obj.DiffPercentForStrike
}

// MaxStrikes returns the max amount of strikes
func (obj *information) MaxStrikes() int {
	return obj.MxStrikes
}
