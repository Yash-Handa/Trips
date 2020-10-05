package driver_test

import (
	"testing"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/db/driver"
)

func TestDriverByID(t *testing.T) {
	// setting up DB
	db.Connect()
	defer db.Close()

	// setting up driver repo
	dr := &driver.Repo{DB: db.GetDB()}
	driverID := "82bdbe64-3d32-43a0-a707-e206c0af5bba"
	want := driver.Driver{
		ID:        driverID,
		FirstName: "Vagita",
		LastName:  nil,
		PhoneNo:   "11111122",
		Gender:    "M",
		Rating:    4.9,
		Pic:       "/raw/assets/drivers/test_Vagita.jpg",
	}

	got, err := dr.GetDriverByID(driverID)
	if err != nil {
		t.Fatalf("Cannot get driver info: %v", err)
	}

	if want != *got {
		t.Fatalf("Required %#v but got %#v", want, *got)
	}
}
