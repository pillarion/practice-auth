package postgresql

import (
	"database/sql"
	"time"
)

// UserDTO defines the user model for the database
type UserDTO struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
