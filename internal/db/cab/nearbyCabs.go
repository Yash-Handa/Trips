package cab

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

// NearbyCabsSub returns a channel which sends nearby cabs
func (cr *Repo) NearbyCabsSub(ctx context.Context, input model.NearbyCabInput) (chan []*model.NearbyCab, error) {
	ch := make(chan []*model.NearbyCab)
	// idsCh := make(chan []string)

	go func(ch chan []*model.NearbyCab, cur *model.LocationInput) {
		prevIDs := make([]string, 0)
		// get slice of ids of available cabs
		IDs, _ := getIDofCabs(input.Type, cr.DB)

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			ch <- createNearbyCab(prevIDs, IDs, cur)
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				prevIDs = IDs[:]
				IDs, _ = getIDofCabs(input.Type, cr.DB)
			}
		}
	}(ch, input.CurLocation)

	return ch, nil
}

func getIDofCabs(cabType model.CabType, db *pg.DB) ([]string, error) {
	var cabs []*Cab

	err := db.Model(&cabs).
		Column("id").
		Where("type = ?", cabType).
		Where("available=true").
		Select()

	if err != nil {
		return nil, fmt.Errorf("Something went wrong")
	}

	str := make([]string, len(cabs))

	for i, v := range cabs {
		str[i] = v.ID
	}

	return str, nil
}

func createNearbyCab(prev, id []string, cur *model.LocationInput) []*model.NearbyCab {
	rlt := make([]*model.NearbyCab, 0)
	for _, i := range id {
		found := false
		for _, j := range prev {
			if i == j {
				rlt = append(rlt, &model.NearbyCab{
					Event:    "MOVE",
					CabID:    i,
					Location: randomLocation(cur),
				})
				found = true
				break
			}
		}
		if found == false {
			rlt = append(rlt, &model.NearbyCab{
				Event:    "ENTER",
				CabID:    i,
				Location: randomLocation(cur),
			})
		}
	}

	for _, i := range prev {
		found := false
		for _, j := range id {
			if i == j {
				found = true
				break
			}
		}
		if found == false {
			rlt = append(rlt, &model.NearbyCab{
				Event:    "LEAVE",
				CabID:    i,
				Location: randomLocation(cur),
			})
		}
	}
	return rlt
}

func randomLocation(cur *model.LocationInput) *model.Location {
	rand.Seed(time.Now().UnixNano())
	return &model.Location{
		Lat: cur.Lat + (-0.00005 + rand.Float64()*0.0001),
		Lon: cur.Lon + (-0.00005 + rand.Float64()*0.0001),
	}
}
