package country

import uuid "github.com/satori/go.uuid"

type country struct {
	UUID *uuid.UUID `json:"id"`
	Nme  string     `json:"name"`
}

func createCountry(id *uuid.UUID, name string) (Country, error) {
	out := country{
		UUID: id,
		Nme:  name,
	}

	return &out, nil
}

func createCountryFromStorable(storable *storableCountry) (Country, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createCountry(&id, storable.Name)
}

// ID returns the ID
func (obj *country) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *country) Name() string {
	return obj.Nme
}
