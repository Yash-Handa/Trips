package user

import "github.com/go-pg/pg/v10"

// User model
type User struct {
	ID        string   `json:"id" pg:"type:uuid,default:gen_random_uuid()"`
	Email     string   `json:"email" pg:",notnull"`
	FirstName string   `json:"firstName" pg:",notnull"`
	LastName  *string  `json:"lastName"`
	TripsID   []string `json:"trips" pg:",array"`
}

// Repo contains all User related functions
type Repo struct {
	DB *pg.DB
}

// DummyUsers creates dummy drivers
func DummyUsers(db *pg.DB) {
	count, err := db.Model(&User{}).Count()
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	l:= "Handa"
	users :=[]User {
		{
			ID: "11111111-1111-1111-1111-111111111111",
			Email: "y@y.com",
			FirstName: "Yash",
			LastName: &l,
		},
		{
			ID: "22222222-2222-2222-2222-222222222222",
			Email: "p@p.com",
			FirstName: "Pooja",
			LastName: &l,
		},
	}

	_, err = db.Model(&users).Insert()
	if err != nil {
		panic(err)
	}
}
