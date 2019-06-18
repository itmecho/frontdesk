package types

import (
	"time"

	"github.com/lib/pq"
)

// User represents a user in the store
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	password  []byte
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt pq.NullTime `json:"deleted_at"`
}
