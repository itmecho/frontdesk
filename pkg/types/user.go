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

func (u *User) SetPassword(pwd string) error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(pwd), 0)
	if err != nil {
		return err
	}

	u.Password = encrypted
	return nil
}
