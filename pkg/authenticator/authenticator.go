package authenticator

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/itmecho/frontdesk/pkg/types"
)

var (
	// ErrInvalidSecret is returned when an empty secret is given
	ErrInvalidSecret = errors.New("Please provide a valid secret")

	// ErrInvalidToken is returned when trying to parse an invalid token
	ErrInvalidToken = errors.New("Token is invalid")
)

// Authenticator handles generating and parsing auth tokens
type Authenticator struct {
	secret []byte
}

// New returns a new instance of and Authenticator
func New(secret []byte) (Authenticator, error) {
	if len(secret) == 0 {
		return Authenticator{}, ErrInvalidSecret
	}
	return Authenticator{secret}, nil
}

type claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// GenerateToken takes a user and generates a JWT containing their user ID
func (a Authenticator) GenerateToken(u *types.User) (string, error) {
	c := claims{
		u.ID,
		jwt.StandardClaims{
			Issuer: "frontdesk",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(a.secret)
}

// CheckToken attempts to parse an auth token
func (a Authenticator) CheckToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.secret, nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}

	return ErrInvalidToken
}
