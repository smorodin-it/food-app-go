package handlers

import (
	_ "food-backend/src/responses"
	"github.com/gofiber/fiber/v2"
	"time"
)

type NestedObject struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type TestDomain struct {
	ID                  string         `json:"id" validate:"required"`
	Required            string         `json:"required" validate:"required"`
	NotRequired         string         `json:"notRequired"`
	Nested              NestedObject   `json:"nestedObject"`
	NestedArray         []NestedObject `json:"nestedArray"`
	NestedArrayOfString []string       `json:"nestedArrayOfString"`
}

type TestForm struct {
	Required      string         `json:"required" validate:"required,min=2,max=100"`
	NotRequired   string         `json:"notRequired" validate:"min=2,max=100"`
	ArrayOfObject []NestedObject `json:"arrayOfObject" validate:"min=1,max=5"`
	ArrayOfString []string       `json:"arrayOfString" validate:"min=2,max=5"`
	NestedObject  *NestedObject  `json:"nestedObject"`
	TestID        string         `json:"testId" validate:"required"`
	DateWhen      time.Time      `json:"dateWhen" validate:"required"`
	OtherDate     time.Time      `json:"otherDate"`
}

type TestHandler interface {
	List() fiber.Handler
	Create() fiber.Handler
	Retrieve() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
	ListOfString() fiber.Handler
}

type testHandler struct {
	s any
}

// List is a function to get list of all test
// @Summary Get test list
// @Description Get test list
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []TestDomain ]
// @Router /test_handler [get]
func (t testHandler) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}

}

// Create is a function to create new test
// @Summary Create new test
// @Description Create new test
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request body TestForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /test_handler [post]
func (t testHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}
}

// Retrieve is a function to get test by id
// @Summary Retrieve test by id
// @Description Retrieve test by id
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ TestDomain ]
// @Router /test_handler/{id} [get]
func (t testHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}
}

// Update is a function to update test by id
// @Summary Update test
// @Description Update test
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Param request body TestForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /test_handler/{id} [put]
func (t testHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}
}

// Delete is a function to delete test
// @Summary Delete test
// @Description Delete test
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request path string true "ingredient in test id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /test_handler/{id} [delete]
func (t testHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}
}

// ListOfString is a function to get list of strings
// @Summary Get test list
// @Description Get test list
// @Tags TestDomain
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []string ]
// @Router /test_handler/strings [get]
func (t testHandler) ListOfString() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//TODO implement me
		panic("implement me")
	}
}

func NewTestHandler(s any) TestHandler {
	return &testHandler{
		s: s,
	}
}
