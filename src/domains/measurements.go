package domains

import "time"

type Measurement struct {
	MeasurementId     string    `json:"measurementId" db:"measurement_id"`
	UserId            string    `json:"userId" db:"user_id"`
	MeasurementWeight int       `json:"measurementWeight" db:"measurement_weight"`
	Date              time.Time `json:"date" db:"date"`
}
