package db

import (
	"time"

	"github.com/Yash-Handa/Trips/internal/gql/model"
)

// Trip type
type Trip struct {
	ID          int               `json:"id" pg:"type:uuid"`
	Pickup      *model.Location   `json:"pickup"`
	Destination *model.Location   `json:"destination"`
	CabID       int               `json:"cab"`
	Amount      *model.Cash       `json:"amount"`
	StartTime   *time.Time        `json:"startTime"`
	EndTime     *time.Time        `json:"duration"`
	Canceled    *model.CancelTrip `json:"canceled"`
	Completed   bool              `json:"completed"`
}
