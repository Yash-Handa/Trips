package trip

import (
	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CancelTrip cancels a non started trip
func (tr *Repo) CancelTrip(ID, uid, reason string) (*Trip, error) {
	t := new(Trip)

	err := tr.DB.Model(t).
		Where("id = ?", ID).
		For("UPDATE").
		Select()

	if err != nil || t == nil {
		return nil, gqlerror.Errorf("No Such trip exist")
	}

	if t.UserID != uid {
		return nil, gqlerror.Errorf("Unauthorized Access")
	}

	if t.Completed == true {
		return nil, gqlerror.Errorf("The trip with ID %s has already been completed", ID)
	}

	if t.StartTime != nil {
		return nil, gqlerror.Errorf("The Trip has already been started")
	}

	t.Completed = true
	t.Canceled = &model.CancelTrip{
		Cancel: true,
		Reason: reason,
	}

	// cancel the trip
	_, err = tr.DB.Model(t).
		Column("completed").
		Column("canceled").
		Where("id = ?", t.ID).
		Update()

	if err != nil {
		return nil, gqlerror.Errorf("Something went wrong")
	}

	// free up the cab/ driver
	_, err = tr.DB.Model(&cab.Cab{Available: true}).
		Column("available").
		Where("id = ?", t.CabID).
		Update()

	if err != nil {
		return nil, gqlerror.Errorf("Something went wrong")
	}

	return t, nil
}
