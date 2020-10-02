package trip

import (
	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// BookTrip books a trip
func (tr *Repo) BookTrip(input model.BookTripInput) (*Trip, error) {
	t := new(Trip)
	t.ID = utils.GenUUID()

	// find a suitable cab
	c := new(cab.Cab)
	err := tr.DB.Model(c).
		Where("type=?", input.CabType).
		Column("id").
		Where("available=true").
		Limit(1).
		For("UPDATE").
		Select()

	if err != nil || c == nil {
		return nil, gqlerror.Errorf("No %s cab is available for your location", input.CabType)
	}

	// mark the cab as non-available
	_, err = tr.DB.Model(&cab.Cab{Available: false}).
		Column("available").
		Where("id = ?", c.ID).
		Update()

	if err != nil || c == nil {
		return nil, gqlerror.Errorf("No %s cab is available for your location", input.CabType)
	}

	t.CabID = c.ID

	// Todo - add user authentication to get user from context
	t.UserID = "11111111-1111-1111-1111-111111111111"

	t.Amount = &model.Cash{
		Amount:   500.00,
		Currency: model.Currency("INR"),
	}

	t.Pickup = &model.Location{
		Lat: input.Pickup.Lat,
		Lon: input.Pickup.Lon,
	}

	t.Destination = &model.Location{
		Lat: input.Destination.Lat,
		Lon: input.Destination.Lon,
	}

	t.Completed = false

	_, err = tr.DB.Model(t).Insert()
	if err != nil {
		return nil, gqlerror.Errorf("Something went Wrong")
	}

	return t, nil
}
