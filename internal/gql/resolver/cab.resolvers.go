package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
)

func (r *cabResolver) Driver(ctx context.Context, obj *db.Cab) (*db.Driver, error) {
	// d := new(model.Driver)

	// for _, v := range r.drivers {
	// 	if v.ID == obj.DriverID {
	// 		d = v
	// 		break
	// 	}
	// }

	// if d == nil {
	// 	return nil, gqlerror.Errorf("%d driver not found", obj.DriverID)
	// }

	// return d, nil
	panic(fmt.Errorf("not implemented"))
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver { return &cabResolver{r} }

type cabResolver struct{ *Resolver }
