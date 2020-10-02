package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) BookTrip(ctx context.Context, input model.BookTripInput) (*model.Trip, error) {
	// a temporary data filler (drivers and cars)
	r.FillDB()

	t := new(model.Trip)
	t.ID = 1234

	// find a suitable cab
	for _, v := range r.cabs {
		if v.Type == input.CabType && v.Available {
			t.CabID = v.ID
			v.Available = false
			break
		}
	}

	if t.CabID == 0 {
		return nil, gqlerror.Errorf("No %s cab is available for your location", input.CabType)
	}

	t.Amount = &model.Cash{
		Amount:   500.00,
		Currency: model.Currency("INR"),
	}

	t.Pickup = &model.Location{
		Lat: input.Pickup.Lat,
		Lon: input.Pickup.Lon,
	}

	t.Destination = &model.Location{
		Lat: input.Destination.Lat,
		Lon: input.Destination.Lon,
	}

	t.Completed = false

	r.trip = append(r.trip, t)
	return t, nil
}

func (r *mutationResolver) CancelTrip(ctx context.Context, id int, reason string) (*model.Trip, error) {
	// a temporary data filler (drivers and cars)
	r.FillDB()

	t := new(model.Trip)
	// find the trip data
	found := false
	for _, v := range r.trip {
		if v.ID == id && v.Completed == false {
			v.Completed = true
			t = v
			found = true
			break
		}
	}

	if found == false {
		return nil, gqlerror.Errorf("Ether the ID %d is wrong or the trip has been completed", id)
	}

	t.Canceled = &model.CancelTrip{
		Cancel: true,
		Reason: reason,
	}

	// update the cab to free it
	for _, v := range r.cabs {
		if v.ID == t.CabID {
			v.Available = true
			break
		}
	}

	return t, nil
}

func (r *mutationResolver) StartTrip(ctx context.Context, id int) (*model.Trip, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EndTrip(ctx context.Context, id int) (*model.Trip, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Trips(ctx context.Context) ([]*model.Trip, error) {
	return r.trip, nil
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
