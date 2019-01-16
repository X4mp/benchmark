package spot

import uuid "github.com/satori/go.uuid"

type spot struct {
	UUID *uuid.UUID `json:"id"`
	Long int        `json:"longitude"`
	Lat  int        `json:"latitude"`
	Rad  int        `json:"radius"`
}

func createSpot(id *uuid.UUID, long int, lat int, rad int) (Spot, error) {
	out := spot{
		UUID: id,
		Long: long,
		Lat:  lat,
		Rad:  rad,
	}

	return &out, nil
}

func createSpotFromStorable(storable *storableSpot) (Spot, error) {
	id, idErr := uuid.FromString(storable.ID)
	if idErr != nil {
		return nil, idErr
	}

	return createSpot(&id, storable.Longitude, storable.Latitude, storable.Radius)
}

// ID returns the ID
func (obj *spot) ID() *uuid.UUID {
	return obj.UUID
}

// Longitude returns the longitude
func (obj *spot) Longitude() int {
	return obj.Long
}

// Latitude returns the latitude
func (obj *spot) Latitude() int {
	return obj.Lat
}

// Radius returns the radius
func (obj *spot) Radius() int {
	return obj.Rad
}
