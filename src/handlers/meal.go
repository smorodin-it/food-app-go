package handlers

import (
	"errors"
	"food-backend/src/forms"
	_ "food-backend/src/responses"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

type MealHandler interface {
	List() fiber.Handler
	ListByUser() fiber.Handler
	Create() fiber.Handler
	Retrieve() fiber.Handler
	Update() fiber.Handler
	AddIngredient() fiber.Handler
	ListIngredient() fiber.Handler
	UpdateIngredient() fiber.Handler
	DeleteIngredient() fiber.Handler
}

type mealHandler struct {
	ms services.MealService
	as services.AuthService
}

// List is a function to get list of all meals
// @Summary Get meals list
// @Description Get meals list
// @Tags Meal
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []domains.Meal ]
// @Router /meal/all [get]
func (h mealHandler) List() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		meals, err := h.ms.List(q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, meals)
	}
}

// ListByUser is a function to get list of all meals by auth user id
// @Summary Get meals list by auth user id
// @Description Get meals list by auth user id
// @Tags Meal
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []domains.Meal ]
// @Router /meal [get]
func (h mealHandler) ListByUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := utils.GetUserIDFromToken(ctx)

		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		meals, err := h.ms.ListByUser(userId, q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, meals)
	}
}

// Create is a function to create new meal
// @Summary Create new meal
// @Description Create new meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request body forms.MealForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /meal [post]
func (h mealHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MealForm)

		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		userId, err := h.as.GetUserIdFromToken(ctx)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := h.ms.Create(*f, userId.(string))
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)

		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// Retrieve is a function to get meal by id
// @Summary Retrieve meal by id
// @Description Retrieve meal by id
// @Tags Meal
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ domains.Meal ]
// @Router /meal/{id} [get]
func (h mealHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		meal, err := h.ms.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, meal)
	}
}

// Update is a function to update meal by id
// @Summary Update meal
// @Description Update meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Param request body forms.MealForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /meal/{id} [put]
func (h mealHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		f := new(forms.MealForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		err = h.ms.Update(*f, id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

// AddIngredient is a function to add ingredient to meal
// @Summary Add ingredient to meal
// @Description Add ingredient to meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request body forms.MealIngredientAddForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /meal/ingredient [post]
func (h mealHandler) AddIngredient() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MealIngredientAddForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := h.ms.AddIngredient(*f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// ListIngredient is a function to get list of all ingredients in meal
// @Summary Get ingredients list in meal
// @Description Get ingredients list in meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request path string true "meal id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []responses.MealIngredientResp ]
// @Router /meal/{id}/ingredient [get]
func (h mealHandler) ListIngredient() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		ingr, err := h.ms.ListIngredients(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, ingr)
	}
}

// UpdateIngredient is a function to update weight of ingredient in meal
// @Summary Update weight of ingredient in meal
// @Description Update weight of ingredient in meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request path string true "ingredient in meal id"
// @Param request body forms.MealIngredientUpdateForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /meal/ingredient/{id} [put]
func (h mealHandler) UpdateIngredient() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		f := new(forms.MealIngredientUpdateForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		err = h.ms.UpdateIngredient(id, *f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

// DeleteIngredient is a function to delete ingredient from meal
// @Summary Delete ingredient from meal
// @Description Delete ingredient from meal
// @Tags Meal
// @Security ApiKeyAuth
// @Param request path string true "ingredient in meal id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /meal/ingredient/{id} [delete]
func (h mealHandler) DeleteIngredient() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		err := h.ms.DeleteIngredient(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

func NewMealHandler(as services.AuthService, ms services.MealService) MealHandler {
	return &mealHandler{
		ms: ms,
		as: as,
	}
}
