package domains

import "time"

type Meal struct {
	ID          string    `json:"id" db:"id"`
	UserId      string    `json:"-" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	TotalWeight int       `json:"totalWeight" db:"total_weight"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
}
