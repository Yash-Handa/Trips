package driver

import (
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

// Driver model
type Driver struct {
	ID        string       `json:"id" pg:"type:uuid,default:gen_random_uuid()"`
	FirstName string       `json:"firstName" pg:",notnull"`
	LastName  *string      `json:"lastName"`
	PhoneNo   string       `json:"phoneNo" pg:",unique,notnull"`
	Gender    model.Gender `json:"gender" pg:",notnull"`
	Rating    float64      `json:"rating" pg:",notnull"`
	Pic       string       `json:"pic" pg:",notnull"`
}

// Repo contains all Driver related functions
type Repo struct {
	DB *pg.DB
}

// DummyDrivers creates dummy drivers
func DummyDrivers(db *pg.DB) {

	count, err := db.Model(&Driver{}).Count()
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	drivers := []Driver{
		{
			ID:        "acb336c4-15f2-47f6-8080-d744173a1f3e",
			FirstName: "Goku",
			PhoneNo:   "11111111",
			Gender:    "M",
			Rating:    4.9,
			Pic:       "test_Goku",
		},
		{
			ID:        "82bdbe64-3d32-43a0-a707-e206c0af5bba",
			FirstName: "Vagita",
			PhoneNo:   "11111122",
			Gender:    "M",
			Rating:    4.9,
			Pic:       "test_Vagita",
		},
		{
			ID:        "388de7d9-e7f1-4338-89e7-2792cf2a4f5f",
			FirstName: "Gohan",
			PhoneNo:   "11111133",
			Gender:    "M",
			Rating:    4.7,
			Pic:       "test_Gohan",
		},
		{
			ID:        "db1eaa9b-753e-4631-84d4-0388b24839bd",
			FirstName: "Trunks",
			PhoneNo:   "11111144",
			Gender:    "M",
			Rating:    4.7,
			Pic:       "test_Trunks",
		},
		{
			ID:        "e69aae35-fd91-45e2-ad70-75df2d96f402",
			FirstName: "Bulma",
			PhoneNo:   "11111155",
			Gender:    "F",
			Rating:    4.7,
			Pic:       "test_Bulma",
		},
		{
			ID:        "db1eaa9b-753e-4631-84d4-0388b24835bd",
			FirstName: "Chi-Chi",
			PhoneNo:   "11111166",
			Gender:    "F",
			Rating:    4.7,
			Pic:       "test_Chi-Chi",
		},
		{
			ID:        "8ba27315-262b-4aa8-80e9-2c984ba9c1f6",
			FirstName: "Frieza",
			PhoneNo:   "11111177",
			Gender:    "M",
			Rating:    4.2,
			Pic:       "test_Frieza",
		},
		{
			ID:        "15990d40-3ae3-4e94-a8fb-a3fafececacb",
			FirstName: "Piccolo",
			PhoneNo:   "11111188",
			Gender:    "M",
			Rating:    4.6,
			Pic:       "test_Piccolo",
		},
		{
			ID:        "c276bf74-cedc-488c-960b-f5214d60188e",
			FirstName: "Beerus",
			PhoneNo:   "11111199",
			Gender:    "M",
			Rating:    5.0,
			Pic:       "test_Beerus",
		},
		{
			ID:        "b1be7dd9-4a3b-4a96-ab2b-cfc662c3e5b0",
			FirstName: "Whis",
			PhoneNo:   "11111100",
			Gender:    "M",
			Rating:    5.0,
			Pic:       "test_Whis",
		},
	}
	_, err = db.Model(&drivers).Insert()
	if err != nil {
		panic(err)
	}
}
