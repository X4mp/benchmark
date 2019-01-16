package city

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/region"
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
	// dependencies:
	country.Register(codec)
	region.Register(codec)

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
