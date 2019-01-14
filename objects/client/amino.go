package client

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
)

const (
	xmnClient           = "xmnnetwork/benchmark/Client"
	xmnNormalizedClient = "xmnnetwork/benchmark/Normalized/Client"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	wallet.Register(codec)
	host.Register(codec)

	// Client
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Client)(nil), nil)
		codec.RegisterConcrete(&client{}, xmnClient, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableClient{}, xmnNormalizedClient, nil)
	}()
}
