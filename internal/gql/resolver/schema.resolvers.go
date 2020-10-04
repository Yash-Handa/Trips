package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Yash-Handa/Trips/internal/db/trip"
	"github.com/Yash-Handa/Trips/internal/db/user"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/Yash-Handa/Trips/internal/middlewares"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var unAuthorized = gqlerror.Errorf("Unauthorized Access")

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*user.AuthResponse, error) {
	return r.UsersRepo.RegisterUser(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*user.AuthResponse, error) {
	return r.UsersRepo.LoginUser(input)
}

func (r *mutationResolver) BookTrip(ctx context.Context, input model.BookTripInput) (*trip.Trip, error) {
	u, err := middlewares.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, unAuthorized
	}

	return r.TripsRepo.BookTrip(u.ID, input)
}

func (r *mutationResolver) CancelTrip(ctx context.Context, id string, reason string) (*trip.Trip, error) {
	u, err := middlewares.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, unAuthorized
	}

	return r.TripsRepo.CancelTrip(id, u.ID, reason)
}

func (r *mutationResolver) StartTrip(ctx context.Context, id string) (*trip.Trip, error) {
	u, err := middlewares.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, unAuthorized
	}

	return r.TripsRepo.StartTrip(id, u.ID)
}

func (r *mutationResolver) EndTrip(ctx context.Context, id string) (*trip.Trip, error) {
	u, err := middlewares.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, unAuthorized
	}

	return r.TripsRepo.EndTrip(id, u.ID)
}

func (r *queryResolver) Trips(ctx context.Context, status model.TripsInput) ([]*trip.Trip, error) {
	u, err := middlewares.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, unAuthorized
	}

	return r.TripsRepo.GetTripsByUser(u.ID, status)
}

func (r *subscriptionResolver) NearbyCabs(ctx context.Context, input model.NearbyCabInput) (<-chan []*model.NearbyCab, error) {
	// this is just a dummy representation
	// Basic logic:
	// Initially all available cabs are given location in proximity of the curLocation
	// These cabs info is send to front-end as event: ENTER
	//
	// then every 5 sec db is queried to check available cabs
	// then some of those cabs (randomly) change there location
	// those cabs info is send to frontend as event: MOVE
	return r.CabsRepo.NearbyCabsSub(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
