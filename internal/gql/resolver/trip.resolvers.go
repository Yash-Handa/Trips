package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
)

func (r *tripResolver) Cab(ctx context.Context, obj *db.Trip) (*db.Cab, error) {
	// c := new(model.Cab)

	// // retrive cab from CabID
	// for _, v := range r.cabs {
	// 	if v.ID == obj.CabID {
	// 		c = v
	// 		break
	// 	}
	// }
	// if c == nil {
	// 	return nil, gqlerror.Errorf("%d cab is not available", obj.CabID)
	// }
	// return c, nil
	panic(fmt.Errorf("not implemented"))
}

// Trip returns generated.TripResolver implementation.
func (r *Resolver) Trip() generated.TripResolver { return &tripResolver{r} }

type tripResolver struct{ *Resolver }
