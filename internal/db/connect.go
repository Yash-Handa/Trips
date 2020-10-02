package db

import (
	"context"

	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/go-pg/pg/v10"
)

var (
	host   string
	port   string
	user   string
	pass   string
	dbName string
)

func init() {
	host = utils.MustGet("POSTGRES_HOST")
	port = utils.MustGet("POSTGRES_PORT")
	user = utils.MustGet("POSTGRES_USER")
	pass = utils.MustGet("POSTGRES_PASSWORD")
	dbName = utils.MustGet("POSTGRES_USER")
}

var db *pg.DB

// Connect to the database
func Connect() {
	db = pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: pass,
		Database: dbName,
	})

	// check if the db is up
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}

	if err := createSchema(db); err != nil {
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
