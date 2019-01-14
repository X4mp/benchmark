package region

import uuid "github.com/satori/go.uuid"

type region struct {
	UUID *uuid.UUID `json:"id"`
	Nme  string     `json:"name"`
}

func createRegion(id *uuid.UUID, name string) (Region, error) {
	out := region{
		UUID: id,
		Nme:  name,
	}

	return &out, nil
}

func createRegionFromStorable(storable *storableRegion) (Region, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createRegion(&id, storable.Name)
}

// ID returns the ID
func (obj *region) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *region) Name() string {
	return obj.Nme
}
