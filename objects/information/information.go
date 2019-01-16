package information

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

type information struct {
	UUID                 *uuid.UUID    `json:"id"`
	NetWallet            wallet.Wallet `json:"network_wallet"`
	MxReqRadius          int           `json:"maximum_request_radius"`
	MinReqInt            time.Duration `json:"minimum_request_interval"`
	Price                int           `json:"price_per_report_purchase"`
	Reward               int           `json:"reward_per_report"`
	MaxSpeedDiff         int           `json:"max_speed_difference_for_noise"`
	DiffPercentForStrike int           `json:"difference_for_strike"`
	MxStrikes            int           `json:"max_strikes"`
}

func createInformation(id *uuid.UUID, netWallet wallet.Wallet, mxRequestRadius int, minReqInt time.Duration, price int, reward int, maxSpeedDiff int, diffPercentForStrike int, maxStrikes int) (Information, error) {
	out := information{
		UUID:                 id,
		NetWallet:            netWallet,
		MxReqRadius:          mxRequestRadius,
		MinReqInt:            minReqInt,
		Price:                price,
		Reward:               reward,
		MaxSpeedDiff:         maxSpeedDiff,
		DiffPercentForStrike: diffPercentForStrike,
		MxStrikes:            maxStrikes,
	}

	return &out, nil
}

func createInformationFromNormalized(normalized *normalizedInformation) (Information, error) {
	id, idErr := uuid.FromString(normalized.ID)
	if idErr != nil {
		return nil, idErr
	}

	walIns, walInsErr := wallet.SDKFunc.CreateMetaData().Denormalize()(normalized.NetworkWallet)
	if walInsErr != nil {
		return nil, walInsErr
	}

	if wal, ok := walIns.(wallet.Wallet); ok {
		return createInformation(
			&id,
			wal,
			normalized.MaximumRequestRadius,
			normalized.MinimumRequestInterval,
			normalized.PricePerReportPurchase,
			normalized.RewardPerReport,
			normalized.MaxSpeedDifferentForNoise,
			normalized.DifferencePercentForStrike,
			normalized.MaxStrikes,
		)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Wallet instance", walIns.ID().String())
	return nil, errors.New(str)

}

func createInformationFromStorable(storable *storableInformation, rep entity.Repository) (Information, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	walID, walIDErr := uuid.FromString(storable.NetworkWalletID)
	if walIDErr != nil {
		return nil, walIDErr
	}

	walIns, walInsErr := rep.RetrieveByID(wallet.SDKFunc.CreateMetaData(), &walID)
	if walInsErr != nil {
		return nil, walInsErr
	}

	if wal, ok := walIns.(wallet.Wallet); ok {
		return createInformation(
			&id,
			wal,
			storable.MaximumRequestRadius,
			storable.MinimumRequestInterval,
			storable.PricePerReportPurchase,
			storable.RewardPerReport,
			storable.MaxSpeedDifferentForNoise,
			storable.DifferencePercentForStrike,
			storable.MaxStrikes,
		)
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Wallet instance", walIns.ID().String())
	return nil, errors.New(str)
}

// ID returns the ID
func (obj *information) ID() *uuid.UUID {
	return obj.UUID
}

// NetworkWallet returns the network wallet
func (obj *information) NetworkWallet() wallet.Wallet {
	return obj.NetWallet
}

// MinimumRequestInterval returns the minimum request interval
func (obj *information) MinimumRequestInterval() time.Duration {
	return obj.MinReqInt
}

// MaximumRequestRadius returns the maximum request radius
func (obj *information) MaximumRequestRadius() int {
	return obj.MxReqRadius
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
