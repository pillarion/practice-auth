package user

import "time"

const (
	RoleUnknown = "UNKNOWN"
	RoleUser    = "USER"
	RoleAdmin   = "ADMIN"
)

// User defines the user model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
