package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Yash-Handa/Trips/internal/db/user"
	"github.com/Yash-Handa/Trips/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

// CurrentUserKey is the key in context corresponding to cur user
const CurrentUserKey = "currentUser"

// Auth authorizes a user
func Auth(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := parseToken(c.Request)
		if err != nil {
			c.Next()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.Next()
			return
		}

		ur := &user.Repo{DB: db}
		user, err := ur.GetUserByID(claims["jti"].(string))
		if err != nil {
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), CurrentUserKey, user)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(utils.MustGet("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, fmt.Errorf("parseToken error: %v", err)
}

// GetCurrentUserFromCTX will be used in Resolvers to extract cur user
func GetCurrentUserFromCTX(ctx context.Context) (*user.User, error) {
	errNoUserInContext := fmt.Errorf("no user in context")

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*user.User)
	if !ok || user.ID == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}
