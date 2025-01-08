package domains

import "time"

type Measurement struct {
	MeasurementId     string    `db:"measurement_id" json:"measurementId" validate:"required"`
	UserId            string    `db:"user_id" json:"userId" validate:"required"`
	MeasurementWeight int       `db:"measurement_weight" json:"measurementWeight" validate:"required"`
	CreatedAt         time.Time `db:"created_at" json:"-" validate:"required"`
	UpdatedAt         time.Time `db:"updated_at" json:"-" validate:"required"`
}
