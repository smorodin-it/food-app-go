package domains

import "time"

type Measurement struct {
	MeasurementId     string    `db:"measurement_id" json:"id" validate:"required"`
	UserId            string    `db:"user_id" json:"-" validate:"required"`
	MeasurementWeight int       `db:"measurement_weight" json:"weight" validate:"required"`
	CreatedAt         time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt         time.Time `db:"updated_at" json:"-" validate:"required"`
}
