package information

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

// Information represents the blockchain information
type Information interface {
	ID() *uuid.UUID
	NetworkWallet() wallet.Wallet
	PricePerReportPurchase() int
	RewardPerReport() int
	MaxSpeedDifferentForNoise() int
	DifferencePercentForStrike() int
	MaxStrikes() int
}

// Normalized represents a normalized information
type Normalized interface {
}

// Repository represents the information repository
type Repository interface {
	Retrieve() (Information, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID                         *uuid.UUID
	NetworkWallet              wallet.Wallet
	PricePerReportPurchase     int
	RewardPerReport            int
	MaxSpeedDifferentForNoise  int
	DifferencePercentForStrike int
	MaxStrikes                 int
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Information SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Information
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Information {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createInformation(
			params.ID,
			params.NetworkWallet,
			params.PricePerReportPurchase,
			params.RewardPerReport,
			params.MaxSpeedDifferentForNoise,
			params.DifferencePercentForStrike,
			params.MaxStrikes,
		)

		if outErr != nil {
			panic(outErr)
		}

		return out
	},
	CreateMetaData: func() entity.MetaData {
		return createMetaData()
	},
	CreateRepresentation: func() entity.Representation {
		return representation()
	},
	CreateRepository: func(params CreateRepositoryParams) Repository {
		metaData := createMetaData()
		out := createRepository(metaData, params.EntityRepository)
		return out
	},
}
