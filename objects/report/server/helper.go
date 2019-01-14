package server

import (
	"errors"
	"fmt"

	rep "github.com/xmnnetwork/benchmark/objects/report"
	serv "github.com/xmnnetwork/benchmark/objects/server"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/datastore"
)

func retrieveAllReportKeyname() string {
	return "reports"
}

func retrieveReportByReportKeyname(repo rep.Report) string {
	base := retrieveAllReportKeyname()
	return fmt.Sprintf("%s:by_report_id:%s", base, repo.ID().String())
}

func retrieveReportByServerKeyname(serve serv.Server) string {
	base := retrieveAllReportKeyname()
	return fmt.Sprintf("%s:by_server_id:%s", base, serve.ID().String())
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
					retrieveReportByReportKeyname(rep.Report()),
					retrieveReportByServerKeyname(rep.Server()),
				}, nil
			}

			str := fmt.Sprintf("the entity (ID: %s) is not a valid Report instance", ins.ID().String())
			return nil, errors.New(str)
		},
		OnSave: func(ds datastore.DataStore, ins entity.Entity) error {
			if repo, ok := ins.(Report); ok {
				// crate metadata and representation:
				reportMetaData := rep.SDKFunc.CreateMetaData()
				serverMetaData := serv.SDKFunc.CreateMetaData()

				// create the repository and service:
				entityRepository := entity.SDKFunc.CreateRepository(ds)

				// make sure the report exists:
				_, retReportErr := entityRepository.RetrieveByID(reportMetaData, repo.Report().ID())
				if retReportErr != nil {
					str := fmt.Sprintf("the given Report (ID: %s) contains a Report (ID: %s) that does not exists", repo.ID().String(), repo.Report().ID().String())
					return errors.New(str)
				}

				// make sure the server exists:
				_, retServerErr := entityRepository.RetrieveByID(serverMetaData, repo.Server().ID())
				if retServerErr != nil {
					str := fmt.Sprintf("the given Report (ID: %s) contains a Server (ID: %s) that does not exists", repo.ID().String(), repo.Server().ID().String())
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
