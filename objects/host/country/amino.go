package country

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnCountry           = "xmnnetwork/benchmark/Country"
	xmnNormalizedCountry = "xmnnetwork/benchmark/Normalized/Country"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Country
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Country)(nil), nil)
		codec.RegisterConcrete(&country{}, xmnCountry, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableCountry{}, xmnNormalizedCountry, nil)
	}()
}
