package organization

import uuid "github.com/satori/go.uuid"

type organization struct {
	UUID *uuid.UUID `json:"id"`
	Nme  string     `json:"name"`
}

func createOrganization(id *uuid.UUID, name string) (Organization, error) {
	out := organization{
		UUID: id,
		Nme:  name,
	}

	return &out, nil
}

func createOrganizationFromStorable(storable *storableOrganization) (Organization, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createOrganization(&id, storable.Name)
}

// ID returns the ID
func (obj *organization) ID() *uuid.UUID {
	return obj.UUID
}

// Name returns the name
func (obj *organization) Name() string {
	return obj.Nme
}
