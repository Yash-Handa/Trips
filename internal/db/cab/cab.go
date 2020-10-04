package cab

import (
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

// Cab model
type Cab struct {
	ID        string        `json:"id" pg:"type:uuid,default:gen_random_uuid()"`
	Type      model.CabType `json:"type" pg:",notnull"`
	Model     string        `json:"model" pg:",notnull"`
	WorkingAc bool          `json:"workingAC" pg:",notnull"`
	Number    int           `json:"number" pg:",notnull"`
	NamePlate string        `json:"namePlate" pg:",unique,notnull"`
	Pic       string        `json:"pic" pg:",notnull"`
	DriverID  string        `json:"driver" pg:"type:uuid,unique,notnull"`
	Available bool          `json:"available" pg:",use_zero,notnull"`
}

// Repo contains all Cab related functions
type Repo struct {
	DB *pg.DB
}

// DummyCabs creates dummy cabs
func DummyCabs(db *pg.DB) {

	count, err := db.Model(&Cab{}).Count()
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	cabs := []Cab{
		{
			ID:        "5596de5e-dad9-4fb1-a43c-60110d6bee0e",
			Type:      "SUV",
			Model:     "Honda CR-V",
			WorkingAc: true,
			Number:    4321,
			NamePlate: "MH 01A B 4321",
			Pic:       "test_Honda_CR-V",
			DriverID:  "acb336c4-15f2-47f6-8080-d744173a1f3e",
			Available: true,
		},
		{
			ID:        "f2d848a9-b1af-4533-93fd-2f8730034fe5",
			Type:      "Luxury",
			Model:     "AUDI A4",
			WorkingAc: true,
			Number:    4322,
			NamePlate: "MH 01A B 4322",
			Pic:       "test_AUDI_A4",
			DriverID:  "82bdbe64-3d32-43a0-a707-e206c0af5bba",
			Available: true,
		},
		{
			ID:        "0603e3b0-579a-4dd9-9b52-5881717bda5d",
			Type:      "Micro",
			Model:     "Datsun Go",
			WorkingAc: true,
			Number:    4323,
			NamePlate: "MH 01A B 4323",
			Pic:       "test_Datsun_Go",
			DriverID:  "388de7d9-e7f1-4338-89e7-2792cf2a4f5f",
			Available: true,
		},
		{
			ID:        "fc256692-4b3e-451e-b39d-b8dcd5cc5787",
			Type:      "Micro",
			Model:     "Hyundai Eon",
			WorkingAc: true,
			Number:    4324,
			NamePlate: "MH 01A B 4324",
			Pic:       "test_Hyundai_Eon",
			DriverID:  "db1eaa9b-753e-4631-84d4-0388b24839bd",
			Available: true,
		},
		{
			ID:        "230b6f63-d66d-4f99-8861-a666f842f57e",
			Type:      "Luxury",
			Model:     "BMW X1",
			WorkingAc: true,
			Number:    4325,
			NamePlate: "MH 01A B 4325",
			Pic:       "test_BMW_X1",
			DriverID:  "e69aae35-fd91-45e2-ad70-75df2d96f402",
			Available: true,
		},
		{
			ID:        "c4e2fedc-f772-4449-b162-ae1275b85971",
			Type:      "Micro",
			Model:     "Maruti Suzuki Wagon R",
			WorkingAc: true,
			Number:    4326,
			NamePlate: "MH 01A B 4326",
			Pic:       "test_Maruti_Suzuki_Wagon_R",
			DriverID:  "db1eaa9b-753e-4631-84d4-0388b24835bd",
			Available: true,
		},
		{
			ID:        "b215157b-6fe7-4ced-a6a8-bbd769ab64b7",
			Type:      "Mini",
			Model:     "Toyota Etios",
			WorkingAc: true,
			Number:    4327,
			NamePlate: "MH 01A B 4327",
			Pic:       "test_Toyota_Etios",
			DriverID:  "8ba27315-262b-4aa8-80e9-2c984ba9c1f6",
			Available: true,
		},
		{
			ID:        "69db8c2e-1a62-429d-b9f4-118d301ac394",
			Type:      "Mini",
			Model:     "Honda City",
			WorkingAc: true,
			Number:    4328,
			NamePlate: "MH 01A B 4328",
			Pic:       "test_Honda_City",
			DriverID:  "15990d40-3ae3-4e94-a8fb-a3fafececacb",
			Available: true,
		},
		{
			ID:        "3d5d9307-df4d-4c12-90e6-89a6ffeb8825",
			Type:      "Mini",
			Model:     "Toyota Corolla",
			WorkingAc: true,
			Number:    4329,
			NamePlate: "MH 01A B 4329",
			Pic:       "test_Toyota_Corolla",
			DriverID:  "c276bf74-cedc-488c-960b-f5214d60188e",
			Available: true,
		},
		{
			ID:        "b9ff39a3-4c6b-4db1-940c-0850ba7468cc",
			Type:      "SUV",
			Model:     "Toyota Innova Crysta",
			WorkingAc: true,
			Number:    4320,
			NamePlate: "MH 01A B 4320",
			Pic:       "test_Toyota_Innova_Crysta",
			DriverID:  "b1be7dd9-4a3b-4a96-ab2b-cfc662c3e5b0",
			Available: true,
		},
	}
	_, err = db.Model(&cabs).Insert()
	if err != nil {
		panic(err)
	}
}
