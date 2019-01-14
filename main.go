package main

import (
	"log"
	"os"

	term "github.com/nsf/termbox-go"
	amino "github.com/tendermint/go-amino"
	cliapp "github.com/urfave/cli"
	"github.com/xmnnetwork/benchmark/cli"
	"github.com/xmnnetwork/benchmark/meta"
)

func reset() {
	term.Sync()
}

func main() {

	// register amino:
	cdc := amino.NewCodec()
	cli.Register(cdc)

	// create the meta to generate the request registry:
	meta.SDKFunc.Create()

	app := cliapp.NewApp()
	app.Version = "2019.01.14"
	app.Name = "xmnbenchmark"
	app.Usage = "This is a benchmark blockchain, where server providers and benchmark requesters create benchmark reports and sells them to customers that needs benchmark data."
	app.Commands = []cliapp.Command{
		*cli.SDKFunc.GenerateConfigs(),
		*cli.SDKFunc.Spawn(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
