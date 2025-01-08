package handlers

import (
	"errors"
	"food-backend/src/forms"
	_ "food-backend/src/responses"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

type IngredientHandler interface {
	List() fiber.Handler
	Create() fiber.Handler
	Retrieve() fiber.Handler
	Update() fiber.Handler
}

type ingredientHandler struct {
	is services.IngredientService
	as services.AuthService
}

// List is a function to get list of all ingredients
// @Summary Get ingredients list
// @Description Get ingredients list
// @Tags Ingredient
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []domains.Ingredient ]
// @Router /ingredient [get]
func (h ingredientHandler) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		list, err := h.is.List(q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		return utils.GetResponseData(ctx, list)
	}
}

// Create is a function to create new ingredient
// @Summary Create new ingredient
// @Description Create new ingredient
// @Tags Ingredient
// @Security ApiKeyAuth
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /ingredient [post]
func (h ingredientHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.IngredientForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		userId, err := h.as.GetUserIdFromToken(ctx)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := h.is.Create(f, userId.(string))
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// Retrieve is a function to get ingredient by id
// @Summary Retrieve ingredient by id
// @Description Retrieve ingredient by id
// @Tags Ingredient
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ domains.Ingredient ]
// @Router /ingredient/{id} [get]
func (h ingredientHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		ingredient, err := h.is.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, ingredient)
	}
}

// Update is a function to update ingredient by id
// @Summary Update ingredient
// @Description Update ingredient
// @Tags Ingredient
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /ingredient/{id} [put]
func (h ingredientHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		f := new(forms.IngredientForm)

		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		err = h.is.Update(f, id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)

		}

		return utils.GetResponseStatus(ctx, true)
	}
}

func NewIngredientHandler(as services.AuthService, is services.IngredientService) IngredientHandler {
	return &ingredientHandler{
		is: is,
		as: as,
	}
}
