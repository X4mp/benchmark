package cli

import (
	amino "github.com/tendermint/go-amino"
	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/customer"
	"github.com/xmnnetwork/benchmark/objects/host"
	"github.com/xmnnetwork/benchmark/objects/host/city"
	"github.com/xmnnetwork/benchmark/objects/host/country"
	"github.com/xmnnetwork/benchmark/objects/host/organization"
	"github.com/xmnnetwork/benchmark/objects/host/region"
	"github.com/xmnnetwork/benchmark/objects/information"
	"github.com/xmnnetwork/benchmark/objects/report"
	report_server "github.com/xmnnetwork/benchmark/objects/report/server"
	"github.com/xmnnetwork/benchmark/objects/server"
	"github.com/xmnservices/xmnsuite/blockchains/core"
)

var cdc = amino.NewCodec()

func init() {
	Register(cdc)
}

// Register registers all the interface -> struct to amino
func Register(codec *amino.Codec) {
	// Dependencies
	client.Register(codec)
	customer.Register(codec)
	host.Register(codec)
	city.Register(codec)
	country.Register(codec)
	organization.Register(codec)
	region.Register(codec)
	information.Register(codec)
	report.Register(codec)
	report_server.Register(codec)
	server.Register(codec)
	core.Register(codec)

	cdc = codec
}
