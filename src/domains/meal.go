package domains

import "time"

type Meal struct {
	Id          string    `json:"id" db:"id"`
	UserId      string    `json:"userId" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	TotalWeight int       `json:"totalWeight" db:"total_weight"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
