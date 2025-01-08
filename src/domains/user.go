package domains

import "time"

type User struct {
	UserID       string    `db:"user_id" json:"id"`
	Email        string    `db:"email" json:"email"`
	RefreshToken *string   `db:"refresh_token" json:"-"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"-"`
	UpdatedAt    time.Time `db:"updated_at" json:"-"`
}
