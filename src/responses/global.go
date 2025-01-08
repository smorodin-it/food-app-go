package responses

type ResponseAdd struct {
	Id string `json:"id"`
}

type ResponseStatus struct {
	Status bool `json:"status"`
}

type ResponseApi[T any] struct {
	Data  T       `json:"data"`
	Error *string `json:"error"`
}
