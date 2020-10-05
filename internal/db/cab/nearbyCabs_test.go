package cab_test

import (
	"context"
	"testing"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/gql/model"
)

func TestNearbyCabs(t *testing.T) {
	// setting up DB
	db.Connect()
	defer db.Close()

	// setting up cab repo
	cr := &cab.Repo{DB: db.GetDB()}

	input := model.NearbyCabInput{
		Type:   "SUV",
		Radius: 5,
		CurLocation: &model.LocationInput{
			Lat: 1.11,
			Lon: 1.11,
		},
	}
	ctx := context.Background()
	defer ctx.Done()

	cabIDs := []string{"5596de5e-dad9-4fb1-a43c-60110d6bee0e", "b9ff39a3-4c6b-4db1-940c-0850ba7468cc"}

	ch, err := cr.NearbyCabsSub(ctx, input)
	if err != nil {
		t.Fatalf("Cannot get cab info: %v", err)
	}

	var data []*model.NearbyCab
	for i := 0; i < 3; i++ {
		data = <-ch
		for _, v := range data {
			if v.CabID != cabIDs[0] && v.CabID != cabIDs[1] {
				t.Fatalf("The retrived ID %s is not known", v.CabID)
			}
		}
	}
}
