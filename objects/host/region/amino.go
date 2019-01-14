package region

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnRegion           = "xmnnetwork/benchmark/Region"
	xmnNormalizedRegion = "xmnnetwork/benchmark/Normalized/Region"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Region
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Region)(nil), nil)
		codec.RegisterConcrete(&region{}, xmnRegion, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableRegion{}, xmnNormalizedRegion, nil)
	}()
}
