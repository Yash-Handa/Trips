package trip

import (
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetTripsByUser gives a list of Trips by user ID
func (tr *Repo) GetTripsByUser(ID string, status model.TripsInput) ([]*Trip, error) {
	var trips []*Trip

	var err error

	switch status {
	case "ALL":
		err = tr.DB.Model(&trips).
			Where("user_id = ?", ID).
			Order("start_time DESC").
			Limit(20).
			Select()
	case "CANCELED":
		err = tr.DB.Model(&trips).
			Where("user_id = ?", ID).
			Where("canceled is not null").
			Limit(20).
			Select()
	case "COMPLETED":
		err = tr.DB.Model(&trips).
			Where("user_id = ?", ID).
			Where("end_time is not null").
			Order("start_time DESC").
			Limit(20).
			Select()
	case "ACTIVE":
		err = tr.DB.Model(&trips).
			Where("user_id = ?", ID).
			Where("completed = false").
			Order("start_time DESC").
			Limit(20).
			Select()
	}

	if err != nil {
		return nil, gqlerror.Errorf("Something went wrong")
	}

	return trips, nil
}
