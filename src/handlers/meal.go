package handlers

import (
	"errors"
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

// MealList is a function to get list of all meals
// @Summary Get meals list
// @Description Get meals list
// @Tags Meal
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} []domains.Meal
// @Router /meal/all [get]
func MealList(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		meals, err := ms.List(q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(meals)
	}
}

// MealListByUser is a function to get list of all meals by auth user id
// @Summary Get meals list by auth user id
// @Description Get meals list by auth user id
// @Tags Meal
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} []domains.Meal
// @Router /meal [get]
func MealListByUser(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := utils.GetUserIDFromToken(ctx)

		q := new(forms.PaginationQuery)

		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		meals, err := ms.ListByUser(userId, q.Page, q.PerPage)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(meals)
	}
}

// MealCreate is a function to create new meal
// @Summary Create new meal
// @Description Create new meal
// @Tags Meal
// @Param request body forms.MealForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseAdd
// @Router /meal [post]
func MealCreate(ms services.MealService, as services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MealForm)

		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		userId, err := as.GetUserIdFromToken(ctx)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := ms.Create(*f, userId.(string))
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)

		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// MealRetrieve is a function to get meal by id
// @Summary Retrieve meal by id
// @Description Retrieve meal by id
// @Tags Meal
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} domains.Meal
// @Router /meal/${id} [get]
func MealRetrieve(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		meal, err := ms.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(meal)
	}
}

// MealUpdate is a function to update meal by id
// @Summary Update meal
// @Description Update meal
// @Tags Meal
// @Param request path string true "id"
// @Param request body forms.MealForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseStatus
// @Router /meal/${id} [put]
func MealUpdate(ms services.MealService) fiber.Handler {
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

		err = ms.Update(*f, id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

// MealAddIngredient is a function to add ingredient to meal
// @Summary Add ingredient to meal
// @Description Add ingredient to meal
// @Tags Meal
// @Param request body forms.MealIngredientAddForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseAdd
// @Router /meal/ingredient [post]
func MealAddIngredient(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MealIngredientAddForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := ms.AddIngredient(*f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

// MealListIngredients is a function to get list of all ingredients in meal
// @Summary Get ingredients list in meal
// @Description Get ingredients list in meal
// @Tags Meal
// @Param request path string true "meal id"
// @Produce json
// @Success 200 {object} []responses.MealIngredientResp
// @Router /meal/${id}/ingredient [get]
func MealListIngredients(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		i, err := ms.ListIngredients(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(i)
	}
}

// MealIngredientUpdate is a function to update weight of ingredient in meal
// @Summary Update weight of ingredient in meal
// @Description Update weight of ingredient in meal
// @Tags Meal
// @Param request path string true "ingredient in meal id"
// @Param request body forms.MealIngredientUpdateForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseAdd
// @Router /meal/ingredient/${id} [put]
func MealIngredientUpdate(ms services.MealService) fiber.Handler {
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

		err = ms.UpdateIngredient(id, *f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

// MealIngredientDelete is a function to delete ingredient from meal
// @Summary Delete ingredient from meal
// @Description Delete ingredient from meal
// @Tags Meal
// @Param request path string true "ingredient in meal id"
// @Produce json
// @Success 200 {object} responses.ResponseStatus
// @Router /meal/ingredient/${id} [delete]
func MealIngredientDelete(ms services.MealService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		err := ms.DeleteIngredient(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": true})
	}
}
