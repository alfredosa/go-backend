package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID           uuid.UUID
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

func (u *User) IsAuthenticated() bool {
	return u.IsActive
}

func (u *User) IsAdminUser() bool {
	return u.IsAdmin
}

func (u *User) CreateUser(db *sqlx.DB) {
	db.MustExec("INSERT INTO users (username, password_hash, email, first_name, last_name, is_active, is_admin) VALUES ($1, $2, $3, $4, $5, true, false)", u.Username, u.PasswordHash, u.Email, u.FirstName, u.LastName)
}

func GetUser(db *sqlx.DB, username string) (User, error) {
	var user User
	err := db.Get(user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		return user, err
	}
	return user, nil
}
