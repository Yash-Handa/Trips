package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *cabResolver) Driver(ctx context.Context, obj *model.Cab) (*model.Driver, error) {
	d := new(model.Driver)

	for _, v := range r.drivers {
		if v.ID == obj.DriverID {
			d = v
			break
		}
	}

	if d == nil {
		return nil, gqlerror.Errorf("%d driver not found", obj.DriverID)
	}

	return d, nil
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver { return &cabResolver{r} }

type cabResolver struct{ *Resolver }
