package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Yash-Handa/Trips/internal/gql/generated"
	"github.com/Yash-Handa/Trips/internal/gql/model"
)

func (r *cabResolver) Driver(ctx context.Context, obj *model.Cab) (*model.Driver, error) {
	panic(fmt.Errorf("not implemented"))
}

// Cab returns generated.CabResolver implementation.
func (r *Resolver) Cab() generated.CabResolver { return &cabResolver{r} }

type cabResolver struct{ *Resolver }
