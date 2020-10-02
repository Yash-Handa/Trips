package model

// Cab type
type Cab struct {
	ID        int     `json:"id"`
	Type      CabType `json:"type"`
	Model     string  `json:"model"`
	WorkingAc bool    `json:"workingAC"`
	Number    int     `json:"number"`
	NamePlate string  `json:"namePlate"`
	Pic       string  `json:"pic"`
	DriverID  int     `json:"driver"`
	Available bool    `json:"available"`
}
