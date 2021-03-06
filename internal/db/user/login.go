package user

import (
	"github.com/Yash-Handa/Trips/internal/gql/model"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var errBadCred = gqlerror.Errorf("EmailID Password combination do not exist")

// LoginUser logs in a user
func (ur *Repo) LoginUser(input model.LoginInput) (*AuthResponse, error) {
	// verification of input
	if !utils.EmailFormatCheck(input.Email) {
		return nil, gqlerror.Errorf("%s is not a valid email", input.Email)
	}

	if len(input.Password) < 6 {
		return nil, gqlerror.Errorf("Password should be more than or equal to 6 charetor")
	}

	u, err := getUserByField(ur.DB, "email", input.Email)
	if err != nil {
		return nil, errBadCred
	}

	// check password hash matching
	err = utils.ComparePassword(input.Password, u.Password)
	if err != nil {
		return nil, errBadCred
	}

	// create JWT Token
	accessToken, exp, err := utils.GenJWTToken(u.ID)
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
	ar.User = u
	return ar, nil
}
