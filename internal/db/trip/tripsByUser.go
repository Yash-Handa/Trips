package trip

import "github.com/vektah/gqlparser/v2/gqlerror"

// GetTripsByUser gives a list of Trips by user ID
func (tr *Repo) GetTripsByUser(ID string) ([]*Trip, error) {
	var trips []*Trip
	err := tr.DB.Model(&trips).
		Order("start_time DESC").
		Limit(20).
		Select()

	if err != nil {
		return nil, gqlerror.Errorf("Something went wrong")
	}

	return trips, nil
}
