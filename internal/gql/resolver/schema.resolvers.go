package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/db/trip"
	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
)

func (r *mutationResolver) BookTrip(ctx context.Context, input model.BookTripInput) (*trip.Trip, error) {
	return r.TripsRepo.BookTrip(input)
}

func (r *mutationResolver) CancelTrip(ctx context.Context, id string, reason string) (*trip.Trip, error) {
	return r.TripsRepo.CancelTrip(id, reason)
}

func (r *mutationResolver) StartTrip(ctx context.Context, id string) (*trip.Trip, error) {
	return r.TripsRepo.StartTrip(id)
}

func (r *mutationResolver) EndTrip(ctx context.Context, id string) (*trip.Trip, error) {
	return r.TripsRepo.EndTrip(id)
}

func (r *queryResolver) Trips(ctx context.Context, status model.TripsInput) ([]*trip.Trip, error) {
	// todo add authentication for user ID
	return r.TripsRepo.GetTripsByUser("11111111-1111-1111-1111-111111111111", status)
}

func (r *subscriptionResolver) NearbyCabs(ctx context.Context, input model.NearbyCabInput) (<-chan []*model.NearbyCab, error) {
	panic(fmt.Errorf("not implemented"))
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
