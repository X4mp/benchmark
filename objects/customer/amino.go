package customer

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

const (
	xmnCustomer           = "xmnnetwork/benchmark/Customer"
	xmnNormalizedCustomer = "xmnnetwork/benchmark/Normalized/Customer"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	transfer.Register(codec)
	report.Register(codec)

	// Customer
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Customer)(nil), nil)
		codec.RegisterConcrete(&customer{}, xmnCustomer, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableCustomer{}, xmnNormalizedCustomer, nil)
	}()
}
