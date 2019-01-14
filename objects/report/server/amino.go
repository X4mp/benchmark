package server

import (
	amino "github.com/tendermint/go-amino"
	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
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
	rep.Register(codec)
	serv.Register(codec)

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
