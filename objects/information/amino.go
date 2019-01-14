package information

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnInformation           = "xmnnetwork/benchmark/Information"
	xmnNormalizedInformation = "xmnnetwork/benchmark/Normalized/Information"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Information
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Information)(nil), nil)
		codec.RegisterConcrete(&information{}, xmnInformation, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableInformation{}, xmnNormalizedInformation, nil)
	}()
}
