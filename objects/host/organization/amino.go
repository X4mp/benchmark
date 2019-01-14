package organization

import (
	amino "github.com/tendermint/go-amino"
)

const (
	xmnOrganization           = "xmnnetwork/benchmark/Organization"
	xmnNormalizedOrganization = "xmnnetwork/benchmark/Normalized/Organization"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Organization
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Organization)(nil), nil)
		codec.RegisterConcrete(&organization{}, xmnOrganization, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableOrganization{}, xmnNormalizedOrganization, nil)
	}()
}
