package domains

import "time"

type User struct {
	Id           string    `json:"id"  db:"id"`
	Email        string    `json:"email" db:"email"`
	RefreshToken *it string   `json:"-" db:"refresh_token"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}
