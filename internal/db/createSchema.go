package db

import (
	"fmt"

	"github.com/Yash-Handa/Trips/internal/db/cab"
	"github.com/Yash-Handa/Trips/internal/db/driver"
	"github.com/Yash-Handa/Trips/internal/db/trip"
	"github.com/Yash-Handa/Trips/internal/db/user"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*cab.Cab)(nil),
		(*trip.Trip)(nil),
		(*driver.Driver)(nil),
		(*user.User)(nil),
	}

	for i, model := range models {
		fmt.Println(i)
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			// IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	// add dummy data to drivers and cabs table
	cab.DummyCabs(db)
	driver.DummyDrivers(db)
	return nil
}
