package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
)

func (r *tripResolver) Cab(ctx context.Context, obj *model.Trip) (*model.Cab, error) {
	panic(fmt.Errorf("not implemented"))
}

// Trip returns generated.TripResolver implementation.
func (r *Resolver) Trip() generated.TripResolver { return &tripResolver{r} }

type tripResolver struct{ *Resolver }
