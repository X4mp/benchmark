package spot

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnSpot           = "xmnnetwork/benchmark/Spot"
	xmnNormalizedSpot = "xmnnetwork/benchmark/Normalized/Spot"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Spot
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Spot)(nil), nil)
		codec.RegisterConcrete(&spot{}, xmnSpot, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableSpot{}, xmnNormalizedSpot, nil)
	}()
}
