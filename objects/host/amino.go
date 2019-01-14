package host

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
)

const (
	xmnHost           = "xmnnetwork/benchmark/Host"
	xmnNormalizedHost = "xmnnetwork/benchmark/Normalized/Host"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	city.Register(codec)
	country.Register(codec)
	organization.Register(codec)
	region.Register(codec)

	// Host
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Host)(nil), nil)
		codec.RegisterConcrete(&host{}, xmnHost, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableHost{}, xmnNormalizedHost, nil)
	}()
}
