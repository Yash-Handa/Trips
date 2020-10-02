package db

import (
	"time"

	"github.com/Yash-Handa/Trips/internal/gql/model"
)

// Trip type
type Trip struct {
	ID          string            `json:"id" pg:"type:uuid,default:gen_random_uuid()"`
	Pickup      *model.Location   `json:"pickup" pg:",notnull"`
	Destination *model.Location   `json:"destination" pg:",notnull"`
	CabID       string            `json:"cab" pg:",notnull"`
	UserID      string            `json:"user" pg:"type:uuid,notnull"`
	Amount      *model.Cash       `json:"amount" pg:",notnull"`
	StartTime   *time.Time        `json:"startTime"`
	EndTime     *time.Time        `json:"duration"`
	Canceled    *model.CancelTrip `json:"canceled"`
	Completed   bool              `json:"completed" pg:",notnull"`
}
