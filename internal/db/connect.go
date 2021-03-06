package db

import (
	"context"

	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

// Connect to the database
func Connect() {
	status := utils.MustGet("STATUS")
	opt := new(pg.Options)

	if status == "dev" {
		opt = &pg.Options{
			Addr:     utils.MustGet("POSTGRES_HOST") + ":" + utils.MustGet("POSTGRES_PORT"),
			User:     utils.MustGet("POSTGRES_USER"),
			Password: utils.MustGet("POSTGRES_PASSWORD"),
			Database: utils.MustGet("POSTGRES_USER"),
		}
	} else {
		// status is heroku and DATABASE_URL will be provided
		var err error
		opt, err = pg.ParseURL(utils.MustGet("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
	}

	db = pg.Connect(opt)

	// check if the db is up
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}

	err = createSchema(db)
	if err != nil {
		panic(err)
	}
}

// Close the connection to database
func Close() {
	db.Close()
}

// GetDB grants access to DB
func GetDB() *pg.DB {
	return db
}
