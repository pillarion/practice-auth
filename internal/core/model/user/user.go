package user

import "time"

const (
	// RoleUnknown is the default role
	RoleUnknown = "UNKNOWN"
	// RoleUser is the user role
	RoleUser = "USER"
	// RoleAdmin is the admin role
	RoleAdmin = "ADMIN"
)

// User defines the user model
type User struct {
	Info      Info
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Info defines the user info model
type Info struct {
	ID       int64
	Name     string
	Email    string
	Role     string
	Password string
}
