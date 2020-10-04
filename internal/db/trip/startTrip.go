package trip

import (
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// StartTrip starts a trip at cur time
func (tr *Repo) StartTrip(ID, uid string) (*Trip, error) {
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

	// if the trip has been canceled or completed
	if t.Completed == true {
		return nil, gqlerror.Errorf("The trip with ID %s has already been completed", ID)
	}

	if t.StartTime != nil {
		return nil, gqlerror.Errorf("The Trip has already been started")
	}

	n := time.Now().UTC()
	t.StartTime = &n

	// Start the trip
	_, err = tr.DB.Model(t).
		Column("start_time").
		Where("id = ?", t.ID).
		Update()

	if err != nil {
		return nil, gqlerror.Errorf("Something went wrong")
	}

	return t, nil
}
