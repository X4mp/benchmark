package report

import (
	"errors"
	"fmt"

	"github.com/xmnnetwork/benchmark/objects/client"
	"github.com/xmnnetwork/benchmark/objects/request"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllReportKeyname() string {
	return "reports"
}

func retrieveReportByRequestKeyname(req request.Request) string {
	base := retrieveAllReportKeyname()
	return fmt.Sprintf("%s:by_request_id:%s", base, req.ID().String())
}

func createMetaData() entity.MetaData {
	return entity.SDKFunc.CreateMetaData(entity.CreateMetaDataParams{
		Name: "Report",
		ToEntity: func(rep entity.Repository, data interface{}) (entity.Entity, error) {
			if storable, ok := data.(*storableReport); ok {
				return createReportFromStorable(storable, rep)
			}

			if dataAsBytes, ok := data.([]byte); ok {
				ptr := new(normalizedReport)
				jsErr := cdc.UnmarshalJSON(dataAsBytes, ptr)
				if jsErr != nil {
					return nil, jsErr
				}

				return createReportFromNormalized(ptr)
			}

			str := fmt.Sprintf("the given data does not represent a Report instance: %s", data)
			return nil, errors.New(str)

		},
		Normalize: func(ins entity.Entity) (interface{}, error) {
			if rep, ok := ins.(Report); ok {
				out, outErr := createNormalizedReport(rep)
				if outErr != nil {
					return nil, outErr
				}

				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Report instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Denormalize: func(ins interface{}) (entity.Entity, error) {
			if normalized, ok := ins.(*normalizedReport); ok {
				return createReportFromNormalized(normalized)
			}

			return nil, errors.New("the given normalized instance cannot be converted to a Report instance")
		},
		EmptyStorable:   new(storableReport),
		EmptyNormalized: new(normalizedReport),
	})
}

func representation() entity.Representation {
	return entity.SDKFunc.CreateRepresentation(entity.CreateRepresentationParams{
		Met: createMetaData(),
		ToStorable: func(ins entity.Entity) (interface{}, error) {
			if rep, ok := ins.(Report); ok {
				out := createStorableReport(rep)
				return out, nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Report instance", ins.ID().String())
			return nil, errors.New(str)
		},
		Keynames: func(ins entity.Entity) ([]string, error) {
			if rep, ok := ins.(Report); ok {
				return []string{
					retrieveAllReportKeyname(),
					retrieveReportByRequestKeyname(rep.Request()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Report instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if rep, ok := ins.(Report); ok {
				// crate metadata and representation:
				clientMetaData := client.SDKFunc.CreateMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)

				// make sure the client exists:
				_, retClientErr := entityRepository.RetrieveByID(clientMetaData, rep.Request().ID())
				if retClientErr != nil {
					str := fmt.Sprintf("the given Report (ID: %s) contains a Request (ID: %s) that does not exists", rep.ID().String(), rep.Request().ID().String())
					return errors.New(str)
				}

				// everything is alright:
				return nil
			}

			str := fmt.Sprintf("the given entity (ID: %s) is not a valid Report instance", ins.ID().String())
			return errors.New(str)
		},
	})
}
