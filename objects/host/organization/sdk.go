package organization

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Organization represents a organization
type Organization interface {
	ID() *uuid.UUID
	Name() string
}

// Normalized represents a normalized organization
type Normalized interface {
}

// Repository represents a organization repository
type Repository interface {
	RetrieveByID(id *uuid.UUID) (Organization, error)
	RetrieveByName(name string) (Organization, error)
	RetrieveSet(index int, amount int) (entity.PartialSet, error)
}

// CreateParams represents the Create params
type CreateParams struct {
	ID   *uuid.UUID
	Name string
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Organization SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Organization
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
	CreateRepository     func(params CreateRepositoryParams) Repository
}{
	Create: func(params CreateParams) Organization {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createOrganization(params.ID, params.Name)
		if outErr != nil {
			panic(outErr)
		}

		return out
	},
	CreateMetaData: func() entity.MetaData {
		return createMetaData()
	},
	CreateRepresentation: func() entity.Representation {
		return representation()
	},
	CreateRepository: func(params CreateRepositoryParams) Repository {
		metaData := createMetaData()
		out := createRepository(metaData, params.EntityRepository)
		return out
	},
}
