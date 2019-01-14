package information

import (
	"errors"
	"fmt"

	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
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

// Retrieve retrieves the information instance
func (app *repository) Retrieve() (Information, error) {
	keynames := []string{
		retrieveAllInformationKeyname(),
	}

	ins, insErr := app.entityRepository.RetrieveByIntersectKeynames(app.metaData, keynames)
	if insErr != nil {
		return nil, insErr
	}

	if inf, ok := ins.(Information); ok {
		return inf, nil
	}

	str := fmt.Sprintf("the entity (ID: %s) is not a valid Information instance", ins.ID().String())
	return nil, errors.New(str)
}
