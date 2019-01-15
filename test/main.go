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
