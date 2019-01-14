package commands

import (
	"math"
	"net"

	"github.com/xmnservices/xmnsuite/blockchains/applications"
	"github.com/xmnservices/xmnsuite/configs"
)

const (
	namespace                    = "xmn"
	name                         = "benchmark"
	id                           = "186f6d61-e62b-4e3f-a32a-ee5303c4cd43"
	databaseFilePath             = "db/blockchain/blockchain.db"
	blockchainRootDirectory      = "db/blockchain/files"
	tokenSymbol                  = "XMB"
	tokenName                    = "XMN Benchmark"
	tokenDescription             = "This is a benchmark blockchain, where server providers and benchmark requesters create benchmark reports and sells them to customers that needs benchmark data."
	totalTokenAmount             = math.MaxInt64 - 1
	initialWalletConcensus       = 50
	initialGazPricePerKB         = 1
	initialTokenConcensusNeeded  = 50
	initialMaxAmountOfValidators = 200
	initialNetworkShare          = 5
	initialValidatorShare        = 80
	initialReferralShare         = 15
	initialUserAmountOfShares    = 100
)

var peers = []string{}

// GenerateConfigsParams represents the generate configs params
type GenerateConfigsParams struct {
	Pass        string
	RetypedPass string
	Filename    string
}

// RetrieveGenesisParams retrieve the genesis transaction params
type RetrieveGenesisParams struct {
	Pass     string
	Filename string
	IP       net.IP
	Port     int
}

// SpawnParams represents the spawn params
type SpawnParams struct {
	Pass     string
	Filename string
	Dir      string
	Port     int
}

// SDKFunc represents the commands SDK func
var SDKFunc = struct {
	GenerateConfigs func(params GenerateConfigsParams) configs.Configs
	Spawn           func(params SpawnParams) applications.Node
}{
	GenerateConfigs: func(params GenerateConfigsParams) configs.Configs {
		out, outErr := generateConfigs(params.Pass, params.RetypedPass, params.Filename)
		if outErr != nil {
			panic(outErr)
		}

		return out
	},
	Spawn: func(params SpawnParams) applications.Node {
		out, outErr := spawn(params.Pass, params.Filename, params.Dir, params.Port)
		if outErr != nil {
			panic(outErr)
		}

		return out
	},
}
