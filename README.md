# Benchmark
This is a benchmark blockchain, where server providers and benchmark requesters create benchmark reports and sells them to customers that needs benchmark data.

## Install dependencies
We use the [dep dependency manager](https://github.com/golang/dep) in order to manage dependencies.  Make sure you have it installed on your computer, if you don't visit the link above and follow instructions to install it.  

Then, install the dependencies using this command:

~~~~
dep ensure
~~~~

## Command line application
To run the command line tool, simply execute this command:

~~~~
go run main.go --help
~~~~

This will show you the possible commands and an explanation for each.

### Generate config file
The config file contains an encrypted private key, protected by a password you give to the tool.  To generate a new config file, simply type this command:

~~~~
go run main.go generate --file ./credentials.xmn --pass MYPASS --rpass MYPASS
~~~~

Make sure to replace the password properly.  You can also replace the filename to whatever file you want.

### Spawn the blockchain
To spawn the blockchain, you need to tell the application where is your config file, and what is the password to decrypt your config file.  Here's the command:

~~~~
go run main.go spawn --dir ./db --pass MYPASS --file ./credentials.xmn
~~~~

The --dir parameter represents the folder where the blockchain will save its data.  The --pass represents the password to decrypt the config file and the --file parameter is the path to your encrypted file.

## Golang SDK
There is a built-in SDK in the xmnsuite package that makes it very easy to communicate with a running blockchain.

### Retrieve an object
To retrieve an instance from the blockchain, you need to sign your query using your private key.  Therefore, you need to load your config instance from your encrypted config file, then execute your query.

Here's an example that retrieve the Information instance.  Please note that I ignore all errors to simplify the example.  Make sure to manage errors in your application.

~~~~
package main

import (
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/information"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/tendermint"
	"github.com/xmnservices/xmnsuite/configs"
)

func main() {

	// variables:
	filePath := "./credentials.xmn"
	pass := "MYPASS"
	blockchainHost := "127.0.0.1:26657"

	// create the config repository:
	confRepository := configs.SDKFunc.CreateRepository()

	// retrieve the config:
	conf, _ := confRepository.Retrieve(filePath, pass)

	// create the client:
	client := tendermint.SDKFunc.CreateClient(tendermint.CreateClientParams{
		IPAsString: blockchainHost,
	})

	// create the entity repository:
	entityRepository := entity.SDKFunc.CreateSDKRepository(entity.CreateSDKRepositoryParams{
		PK:     conf.WalletPK(),
		Client: client,
	})

	// create the specific repository of the object you want to retrieve:
	informationRepository := information.SDKFunc.CreateRepository(information.CreateRepositoryParams{
		EntityRepository: entityRepository,
	})

	// then retrieve the instance:
	inf, _ := informationRepository.Retrieve()

	// print:
	fmt.Sprintf("Instance: %v", inf)
}

~~~~
