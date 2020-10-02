package resolver

//go:generate go run github.com/99designs/gqlgen --verbose

// This file will not be regenerated automatically.

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	// drivers []*model.Driver
	// cabs    []*model.Cab
	// trip    []*model.Trip
	// db
}

// FillDB fills temporary data
// func (r *Resolver) FillDB() {
// 	if len(r.drivers) == 0 {
// 		r.drivers = []*model.Driver{
// 			{
// 				ID:        1,
// 				FirstName: "Goku",
// 				PhoneNo:   "11111111",
// 				Gender:    "M",
// 				Rating:    4.5,
// 				Pic:       "test_Goku",
// 			},
// 			{
// 				ID:        2,
// 				FirstName: "Vagita",
// 				PhoneNo:   "11111122",
// 				Gender:    "M",
// 				Rating:    4.5,
// 				Pic:       "test_Vagita"},
// 		}
// 	}

// 	if len(r.cabs) == 0 {
// 		r.cabs = []*model.Cab{
// 			{
// 				ID:        10,
// 				Type:      "SUV",
// 				Model:     "Honda 123",
// 				WorkingAc: true,
// 				Number:    1234,
// 				NamePlate: "DL ACT 1234",
// 				Pic:       "test_Honda_123",
// 				DriverID:  1,
// 				Available: true,
// 			},
// 			{
// 				ID:        20,
// 				Type:      "Luxury",
// 				Model:     "AUDI 123",
// 				WorkingAc: true,
// 				Number:    6789,
// 				NamePlate: "DL ACT 6789",
// 				Pic:       "test_AUDI_123",
// 				DriverID:  2,
// 				Available: true,
// 			},
// 		}
// 	}
// }
