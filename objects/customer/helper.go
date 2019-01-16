package customer

import (
	"errors"
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllCustomerKeyname() string {
	return "customers"
}

func retrieveCustomerByTransferKeyname(trsf transfer.Transfer) string {
	base := retrieveAllCustomerKeyname()
	return fmt.Sprintf("%s:by_transfer_id:%s", base, trsf.ID().String())
}

func retrieveCustomerByReportKeyname(rep report.Report) string {
	base := retrieveAllCustomerKeyname()
	return fmt.Sprintf("%s:by_report_id:%s", base, rep.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Customer",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableCustomer); ok {
				return createCustomerFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedCustomer)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createCustomerFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Customer instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if cus, ok := ins.(Customer); ok {
				out, outErr := createNormalizedCustomer(cus)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Customer instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedCustomer); ok {
				return createCustomerFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Customer instance")
		},
		EmptyStorable:   new(storableCustomer),
		EmptyNormalized: new(normalizedCustomer),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if cus, ok := ins.(Customer); ok {
				out := createStorableCustomer(cus)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Customer instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if cus, ok := ins.(Customer); ok {
				keynames := []string{
					retrieveAllCustomerKeyname(),
					retrieveCustomerByTransferKeyname(cus.Transfer()),
				}

				reps := cus.Reports()
				for _, oneRep := range reps {
					keynames = append(keynames, retrieveCustomerByReportKeyname(oneRep))
				}

				return keynames, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Customer instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if cus, ok := ins.(Customer); ok {
				// crate metadata and representation:
				reportMetaData := report.SDKFunc.CreateMetaData()
				transferRepresentation := transfer.SDKFunc.CreateRepresentation()
				transferMetaData := transferRepresentation.MetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)
				entityService := entity.SDKFunc.CreateService(ds)

				// make sure the reports exists:
				reps := cus.Reports()
				for _, oneRep := range reps {
					_, retRepErr := entityRepository.RetrieveByID(reportMetaData, oneRep.ID())
					if retRepErr != nil {
						str := fmt.Sprintf("the Customer (ID: %s) contains a Report (ID: %s) that does not exists", cus.ID().String(), oneRep.ID().String())
						return errors.New(str)
					}
				}

				// make sure the transfer does not exists:
				_, retTrsfErr := entityRepository.RetrieveByID(transferMetaData, cus.Transfer().ID())
				if retTrsfErr == nil {
					str := fmt.Sprintf("the Customer (ID: %s) contains a Transfer (ID: %s) that already exists", cus.ID().String(), cus.Transfer().ID().String())
					return errors.New(str)
				}

				// save the transfer:
				saveTrsfErr := entityService.Save(cus.Transfer(), transferRepresentation)
				if saveTrsfErr != nil {
					return saveTrsfErr
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Customer instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
