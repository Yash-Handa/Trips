package user

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetUserByID gets user by id
func (ur *Repo) GetUserByID(ID string) (*User, error) {
	u := new(User)
	err := ur.DB.Model(u).Where("id = ?", ID).Select()
	if err != nil {
		return nil, gqlerror.Errorf("User not found")
	}
	return u, nil
}
