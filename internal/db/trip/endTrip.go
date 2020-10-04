package trip

import (
	"time"

	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// EndTrip stops a started a trip
func (tr *Repo) EndTrip(ID, uid string) (*Trip, error) {
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

	if t.StartTime == nil {
		return nil, gqlerror.Errorf("The Trip has not been started yet")
	}

	t.Completed = true
	n := time.Now().UTC()
	t.EndTime = &n

	// cancel the trip
	_, err = tr.DB.Model(t).
		Column("completed").
		Column("end_time").
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
