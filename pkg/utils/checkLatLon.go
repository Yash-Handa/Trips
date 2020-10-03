package utils

import "fmt"

// CheckLatLon checks if latitude and longitude are valid
func CheckLatLon(lat, lon float64) error {
	str := ""
	if lat > 90 || lat < -90 {
		str += fmt.Sprintf("The Lat provided: %v is not valid, should be between -90 and 90.", lat)
	}

	if lon > 180 || lon < -180 {
		str += fmt.Sprintf("The Lon provided: %v is not valid, should be between -180 and 180.", lon)
	}

	if str == "" {
		return nil
	}

	return fmt.Errorf(str)
}
