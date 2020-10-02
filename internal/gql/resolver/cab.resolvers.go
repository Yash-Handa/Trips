package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/db/driver"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
)

func (r *cabResolver) Driver(ctx context.Context, obj *cab.Cab) (*driver.Driver, error) {
	return r.DriversRepo.GetDriverByID(obj.DriverID)
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver { return &cabResolver{r} }

type cabResolver struct{ *Resolver }
