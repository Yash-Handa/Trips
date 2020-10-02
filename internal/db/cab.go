package db

import "github.com/Yash-Handa/Trips/internal/gql/model"

// Cab model
type Cab struct {
	ID        int           `json:"id" pg:"type:uuid"`
	Type      model.CabType `json:"type"`
	Model     string        `json:"model"`
	WorkingAc bool          `json:"workingAC"`
	Number    int           `json:"number"`
	NamePlate string        `json:"namePlate"`
	Pic       string        `json:"pic"`
	DriverID  int           `json:"driver"`
	Available bool          `json:"available"`
}
