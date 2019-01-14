package server

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/pledge"
)

const (
	xmnServer           = "xmnnetwork/benchmark/Server"
	xmnNormalizedServer = "xmnnetwork/benchmark/Normalized/Server"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	host.Register(codec)
	pledge.Register(codec)

	// Server
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Server)(nil), nil)
		codec.RegisterConcrete(&server{}, xmnServer, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableServer{}, xmnNormalizedServer, nil)
	}()
}
