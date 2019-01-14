package city

import uuid "github.com/satori/go.uuid"

type city struct {
	UUID *uuid.UUID `json:"id"`
	Nme  string     `json:"name"`
}

func createCity(id *uuid.UUID, name string) (City, error) {
	out := city{
		UUID: id,
		Nme:  name,
	}

	return &out, nil
}

func createCityFromStorable(storable *storableCity) (City, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createCity(&id, storable.Name)
}

// ID returns the ID
func (obj *city) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *city) Name() string {
	return obj.Nme
}
