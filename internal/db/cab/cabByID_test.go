package cab_test

import (
	"testing"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/db/cab"
)

func TestCabByID(t *testing.T) {
	// setting up DB
	db.Connect()
	defer db.Close()

	// setting up cab repo
	cr := &cab.Repo{DB: db.GetDB()}
	cabID := "5596de5e-dad9-4fb1-a43c-60110d6bee0e"
	want := cab.Cab{
		ID:        cabID,
		Type:      "SUV",
		Model:     "Honda CR-V",
		WorkingAc: true,
		Number:    4321,
		NamePlate: "MH 01A B 4321",
		Pic:       "/raw/assets/cabs/test_Honda_CR-V.jpg",
		DriverID:  "acb336c4-15f2-47f6-8080-d744173a1f3e",
		Available: true,
	}

	got, err := cr.GetCabByID(cabID)
	if err != nil {
		t.Fatalf("Cannot get cab info: %v", err)
	}

	if want != *got {
		t.Fatalf("Required %#v but got %#v", want, *got)
	}
}
