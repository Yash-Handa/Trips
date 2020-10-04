package user

import (
	"fmt"
	"time"

	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/go-pg/pg/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// RegisterUser registers users
func (ur *Repo) RegisterUser(input model.RegisterInput) (*AuthResponse, error) {
	_, err := getUserByField(ur.DB, "email", input.Email)
	if err == nil {
		return nil, gqlerror.Errorf("Email ID; %s has already been taken", input.Email)
	}

	// create user
	newUser := &User{
		ID:        utils.GenUUID(),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CreatedAt: time.Now().UTC(),
	}

	// generate the hash pass
	pass, err := utils.HashPassword(input.Password)
	if err != nil {
		// todo log to server
		return nil, gqlerror.Errorf("something went wrong")
	}

	newUser.Password = pass

	// add user to DB
	_, err = ur.DB.Model(newUser).Insert()
	if err != nil {
		// todo log to server
		return nil, gqlerror.Errorf("Something went Wrong")
	}

	// create JWT Token
	accessToken, exp, err := utils.GenJWTToken(newUser.ID)
	if err != nil {
		// todo log to server
		return nil, gqlerror.Errorf("Something went Wrong")
	}

	// Create the Auth response
	ar := new(AuthResponse)
	ar.AuthToken = &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   exp,
	}
	ar.User = newUser
	return ar, nil
}

func getUserByField(db *pg.DB, field, value string) (*User, error) {
	var u User
	err := db.Model(&u).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &u, err
}
