package customer

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xmnnetwork/benchmark/objects/report"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity/entities/wallet/entities/transfer"
)

type repository struct {
	metaData         entity.MetaData
	entityRepository entity.Repository
}

func createRepository(metaData entity.MetaData, entityRepository entity.Repository) Repository {
	out := repository{
		metaData:         metaData,
		entityRepository: entityRepository,
	}

	return &out
}

// RetrieveByID retrieves the customer by ID
func (app *repository) RetrieveByID(id *uuid.UUID) (Customer, error) {
	ins, insErr := app.entityRepository.RetrieveByID(app.metaData, id)
	if insErr != nil {
		return nil, insErr
	}

	if cus, ok := ins.(Customer); ok {
		return cus, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Customer instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveByTransfer retrieves the customer by transfer
func (app *repository) RetrieveByTransfer(trsf transfer.Transfer) (Customer, error) {
	keynames := []string{
		retrieveAllCustomerKeyname(),
		retrieveCustomerByTransferKeyname(trsf),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if cus, ok := ins.(Customer); ok {
		return cus, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Customer instance", ins.ID().String())
	return nil, errors.New(str)
}

// RetrieveSetByReport retrieves the customer set by report
func (app *repository) RetrieveSetByReport(rep report.Report, index int, amount int) (entity.PartialSet, error) {
	keynames := []string{
		retrieveAllCustomerKeyname(),
		retrieveCustomerByReportKeyname(rep),
	}

	return app.entityRepository.RetrieveSetByIntersectKeynames(app.metaData, keynames, index, amount)
}
