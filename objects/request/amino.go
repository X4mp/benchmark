package request

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/spot"
)

const (
	xmnRequest           = "xmnnetwork/benchmark/Request"
	xmnNormalizedRequest = "xmnnetwork/benchmark/Normalized/Request"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	client.Register(codec)
	city.Register(codec)
	country.Register(codec)
	organization.Register(codec)
	region.Register(codec)
	spot.Register(codec)

	// Request
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Request)(nil), nil)
		codec.RegisterConcrete(&request{}, xmnRequest, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableRequest{}, xmnNormalizedRequest, nil)
	}()
}
