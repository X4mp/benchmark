package city

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnCity           = "xmnnetwork/benchmark/City"
	xmnNormalizedCity = "xmnnetwork/benchmark/Normalized/City"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// City
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*City)(nil), nil)
		codec.RegisterConcrete(&city{}, xmnCity, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableCity{}, xmnNormalizedCity, nil)
	}()
}
