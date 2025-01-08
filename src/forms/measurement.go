package forms

type MeasurementCreateForm struct {
	UserId            string `json:"userId" validate:"required"`
	MeasurementWeight int    `json:"weight" validate:"required"`
}

type MeasurementUpdateForm struct {
	MeasurementWeight int `json:"weight" validate:"required"`
}
