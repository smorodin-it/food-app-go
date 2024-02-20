package handlers

import (
	"errors"
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

// IngredientList is a function to get list of all ingredients
// @Summary Get ingredients list
// @Description Get ingredients list
// @Tags Ingredient
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} []domains.Ingredient
// @Router /ingredient [get]
func IngredientList(is services.IngredientService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		list, err := is.List(q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(list)
	}
}

// IngredientCreate is a function to create new ingredient
// @Summary Create new ingredient
// @Description Create new ingredient
// @Tags Ingredient
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseAdd
// @Router /ingredient [post]
func IngredientCreate(is services.IngredientService, as services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.IngredientForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		userId, err := as.GetUserIdFromToken(ctx)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := is.Create(f, userId.(string))
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// IngredientRetrieve is a function to get ingredient by id
// @Summary Retrieve ingredient by id
// @Description Retrieve ingredient by id
// @Tags Ingredient
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} domains.Ingredient
// @Router /ingredient/${id} [get]
func IngredientRetrieve(is services.IngredientService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		ingredient, err := is.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(ingredient)
	}
}

// IngredientUpdate is a function to update ingredient by id
// @Summary Update ingredient
// @Description Update ingredient
// @Tags Ingredient
// @Param request path string true "id"
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseStatus
// @Router /ingredient/${id} [put]
func IngredientUpdate(is services.IngredientService) fiber.Handler {
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

		err = is.Update(f, id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)

		}

		return utils.GetResponseStatus(ctx, true)
	}
}
