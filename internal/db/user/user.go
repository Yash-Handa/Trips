package user

import (
	"time"

	"github.com/go-pg/pg/v10"
)

// User model
type User struct {
	ID        string    `json:"id" pg:"type:uuid,unique,default:gen_random_uuid()"`
	Email     string    `json:"email" pg:",unique,notnull"`
	Password  string    `json:"password" pg:",notnull"`
	FirstName string    `json:"firstName" pg:",notnull"`
	LastName  *string   `json:"lastName"`
	CreatedAt time.Time `json:"createdAt" pg:"type:timestamp,notnull"`
}

// AuthResponse model
type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

// AuthToken model
type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

// Repo contains all User related functions
type Repo struct {
	DB *pg.DB
}
