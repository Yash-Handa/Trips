package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Cab)(nil),
		(*Trip)(nil),
		(*Driver)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	// add dummy data to drivers and cabs table
	dummyCabs(db)
	dummyDrivers(db)
	return nil
}
