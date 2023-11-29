package models

import (
	"time"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Email        string
	FirstName    string
	LastName     string
	IsActive     bool
	IsAdmin      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Chirp struct {
	ID    int
	Chirp string
}
