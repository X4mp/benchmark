package meta

import (
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
	"github.com/xmnservices/xmnsuite/blockchains/core/meta"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/underlying/token"
)

func createMeta() (meta.Meta, error) {
	// create core metadata:
	walletMetaData := wallet.SDKFunc.CreateMetaData()
	tokenMetaData := token.SDKFunc.CreateMetaData()

	// create the representations:
	informationRepresentation := information.SDKFunc.CreateRepresentation()
	clientRepresentation := client.SDKFunc.CreateRepresentation()
	customerRepresentation := customer.SDKFunc.CreateRepresentation()
	hostRepresentation := host.SDKFunc.CreateRepresentation()
	hostCityRepresentation := city.SDKFunc.CreateRepresentation()
	hostCountryRepresentation := country.SDKFunc.CreateRepresentation()
	hostRegionRepresentation := region.SDKFunc.CreateRepresentation()
	hostOrganizationRepresentation := organization.SDKFunc.CreateRepresentation()
	reportServerRepresentation := report_server.SDKFunc.CreateRepresentation()
	reportRepresentation := report.SDKFunc.CreateRepresentation()
	serverRepresentation := server.SDKFunc.CreateRepresentation()

	// get the metadata:
	informationMetaData := informationRepresentation.MetaData()
	clientMetaData := clientRepresentation.MetaData()
	customerMetaData := customerRepresentation.MetaData()
	hostMetaData := hostRepresentation.MetaData()
	hostCityMetaData := hostCityRepresentation.MetaData()
	hostCountryMetaData := hostCountryRepresentation.MetaData()
	hostRegionMetaData := hostRegionRepresentation.MetaData()
	hostOrganizationMetaData := hostOrganizationRepresentation.MetaData()
	reportServerMetaData := reportServerRepresentation.MetaData()
	reportMetaData := reportRepresentation.MetaData()
	serverMetaData := serverRepresentation.MetaData()

	// create the meta:
	met := meta.SDKFunc.Create(meta.CreateParams{
		AdditionalRead: map[string]entity.MetaData{
			informationMetaData.Keyname():      informationMetaData,
			clientMetaData.Keyname():           clientMetaData,
			customerMetaData.Keyname():         customerMetaData,
			hostMetaData.Keyname():             hostMetaData,
			hostCityMetaData.Keyname():         hostCityMetaData,
			hostCountryMetaData.Keyname():      hostCountryMetaData,
			hostRegionMetaData.Keyname():       hostRegionMetaData,
			hostOrganizationMetaData.Keyname(): hostOrganizationMetaData,
			reportServerMetaData.Keyname():     reportServerMetaData,
			reportMetaData.Keyname():           reportMetaData,
			serverMetaData.Keyname():           serverMetaData,
		},
	})

	// list of entities that are voted by wallet owners:
	walletOwners := []entity.Representation{
		clientRepresentation,
		customerRepresentation,
		reportServerRepresentation,
		reportRepresentation,
		serverRepresentation,
	}

	for _, oneRep := range walletOwners {
		addedToWalletVoteErr := met.AddToWriteOnEntityRequest(walletMetaData, oneRep)
		if addedToWalletVoteErr != nil {
			return nil, addedToWalletVoteErr
		}
	}

	// list of entities that are voted by token owners:
	tokenOwners := []entity.Representation{
		informationRepresentation,
	}

	for _, oneRep := range tokenOwners {
		addedToTokenVoteErr := met.AddToWriteOnEntityRequest(tokenMetaData, oneRep)
		if addedToTokenVoteErr != nil {
			return nil, addedToTokenVoteErr
		}
	}

	// returns:
	return met, nil
}
