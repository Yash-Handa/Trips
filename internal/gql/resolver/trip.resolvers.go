package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/db/trip"
	"github.com/Yash-Handa/Trips/internal/db/user"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
)

func (r *tripResolver) Cab(ctx context.Context, obj *trip.Trip) (*cab.Cab, error) {
	return r.CabsRepo.GetCabByID(obj.CabID)
}

func (r *tripResolver) User(ctx context.Context, obj *trip.Trip) (*user.User, error) {
	return r.UsersRepo.GetUserByID(obj.UserID)
}

// Trip returns generated.TripResolver implementation.
func (r *Resolver) Trip() generated.TripResolver { return &tripResolver{r} }

type tripResolver struct{ *Resolver }
