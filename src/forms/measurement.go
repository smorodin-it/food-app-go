package forms

type MeasurementCreateForm struct {
	UserId            string `json:"userId" validate:"required"`
	MeasurementWeight int    `json:"measurementWeight" validate:"required"`
}

type MeasurementUpdateForm struct {
	MeasurementWeight int `json:"measurementWeight" validate:"required"`
}
