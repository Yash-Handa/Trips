package resolver

import (
	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/db/driver"
	"github.com/Yash-Handa/Trips/internal/db/trip"
	"github.com/Yash-Handa/Trips/internal/db/user"
)

//go:generate go run github.com/99designs/gqlgen --verbose

// This file will not be regenerated automatically.

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	TripsRepo   *trip.Repo
	CabsRepo    *cab.Repo
	DriversRepo *driver.Repo
	UsersRepo   *user.Repo
}
