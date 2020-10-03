package trip

import (
	"time"

	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

// Trip type
type Trip struct {
	ID          string            `json:"id" pg:"type:uuid,default:gen_random_uuid()"`
	Pickup      *model.Location   `json:"pickup" pg:",notnull"`
	Destination *model.Location   `json:"destination" pg:",notnull"`
	CabID       string            `json:"cab" pg:"type:uuid,notnull"`
	UserID      string            `json:"user" pg:"type:uuid,notnull"`
	Amount      *model.Cash       `json:"amount" pg:",notnull"`
	StartTime   *time.Time        `json:"startTime" pg:"type:timestamp"`
	EndTime     *time.Time        `json:"endTime" pg:"type:timestamp"`
	Canceled    *model.CancelTrip `json:"canceled"`
	Completed   bool              `json:"completed" pg:",use_zero,notnull"`
}

// Repo contains all Trips related functions
type Repo struct {
	DB *pg.DB
}
