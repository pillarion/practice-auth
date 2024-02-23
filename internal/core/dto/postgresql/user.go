package postgresql

import (
	"database/sql"
	"time"
)

type UserDTO struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
