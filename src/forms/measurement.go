package forms

type MeasurementCreateForm struct {
	UserId            string `json:"userId"`
	MeasurementWeight int    `json:"measurementWeight"`
}

type MeasurementUpdateForm struct {
	MeasurementWeight int `json:"measurementWeight"`
}
