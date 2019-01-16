package spot

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmnservices/xmnsuite/blockchains/core/objects/entity"
)

// Spot represents a spot
type Spot interface {
	ID() *uuid.UUID
	Longitude() int
	Latitude() int
	Radius() int
}

// Normalized represents a normalized spot
type Normalized interface {
}

// CreateParams represents the Create params
type CreateParams struct {
	ID        *uuid.UUID
	Longitude int
	Latitude  int
	Radius    int
}

// CreateRepositoryParams represents the CreateRepository params
type CreateRepositoryParams struct {
	EntityRepository entity.Repository
}

// SDKFunc represents the Information SDK func
var SDKFunc = struct {
	Create               func(params CreateParams) Spot
	CreateMetaData       func() entity.MetaData
	CreateRepresentation func() entity.Representation
}{
	Create: func(params CreateParams) Spot {
		if params.ID == nil {
			id := uuid.NewV4()
			params.ID = &id
		}

		out, outErr := createSpot(params.ID, params.Longitude, params.Latitude, params.Radius)
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
}
