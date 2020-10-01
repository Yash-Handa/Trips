package model

import "time"

// Trip type
type Trip struct {
	ID          string      `json:"id"`
	Pickup      *Location   `json:"pickup"`
	Destination *Location   `json:"destination"`
	CabID       string      `json:"cab"`
	Amount      *Cash       `json:"amount"`
	StartTime   *time.Time  `json:"startTime"`
	Duration    *time.Time  `json:"duration"`
	Canceled    *CancelTrip `json:"canceled"`
	Completed   bool        `json:"completed"`
}
