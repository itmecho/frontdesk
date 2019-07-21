package types

import (
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the store
type User struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Password  []byte      `json:"-"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt pq.NullTime `json:"-" db:"deleted_at"`
}

// SetPassword takes a plain text password, hashes it and stores it in the password field on the user
func (u *User) SetPassword(pwd []byte) error {
	encrypted, err := bcrypt.GenerateFromPassword(pwd, 0)
	if err != nil {
		return err
	}

	u.Password = encrypted
	return nil
}

// CheckPassword compares a plain text password to the hashed password stored on the user and returns true or false if they match or don't match respectively
func (u *User) CheckPassword(checkPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(u.Password, checkPassword); err != nil {
		return false
	}

	return true
}
