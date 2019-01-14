package report

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/client"
)

const (
	xmnReport           = "xmnnetwork/benchmark/Report"
	xmnNormalizedReport = "xmnnetwork/benchmark/Normalized/Report"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// dependencies:
	client.Register(codec)

	// Report
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Report)(nil), nil)
		codec.RegisterConcrete(&report{}, xmnReport, nil)
	}()

	// Normalized
	func() {
		defer func() {
			recover()
		}()
		codec.RegisterInterface((*Normalized)(nil), nil)
		codec.RegisterConcrete(&storableReport{}, xmnNormalizedReport, nil)
	}()
}
