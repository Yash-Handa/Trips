package model

import "time"

// Trip type
type Trip struct {
	ID          int         `json:"id"`
	Pickup      *Location   `json:"pickup"`
	Destination *Location   `json:"destination"`
	CabID       int         `json:"cab"`
	Amount      *Cash       `json:"amount"`
	StartTime   *time.Time  `json:"startTime"`
	EndTime     *time.Time  `json:"duration"`
	Canceled    *CancelTrip `json:"canceled"`
	Completed   bool        `json:"completed"`
}
