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

// Authenticator describes how an authenticator should operate
type Authenticator interface {
	GenerateToken(*types.User) (string, error)
	CheckToken(string) error
}

// JWTAuthenticator handles generating and parsing JWTs
type JWTAuthenticator struct {
	secret []byte
}

// NewJWTAuthenticator returns a new instance of a JWTAuthenticator
func NewJWTAuthenticator(secret []byte) (Authenticator, error) {
	if len(secret) == 0 {
		return JWTAuthenticator{}, ErrInvalidSecret
	}
	return JWTAuthenticator{secret}, nil
}

type claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// GenerateToken takes a user and generates a JWT containing their user ID
func (a JWTAuthenticator) GenerateToken(u *types.User) (string, error) {
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
func (a JWTAuthenticator) CheckToken(tokenString string) error {
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
