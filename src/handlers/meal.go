package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

func MealList(ctx *fiber.Ctx) error {
	q := new(forms.PaginationQuery)

	err := ctx.QueryParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	meals, err := ms.List(q.Page, q.PerPage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(meals)
}

func MealListByUser(ctx *fiber.Ctx) error {
	userId := utils.GetUserIDFromToken(ctx)

	q := new(forms.PaginationQuery)

	err := ctx.QueryParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	meals, err := ms.ListByUser(userId, q.Page, q.PerPage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(meals)
}

func MealCreate(ctx *fiber.Ctx) error {
	f := new(forms.MealForm)

	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	as := new(services.AuthService)
	userId, err := as.GetUserIdFromToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	id, err := ms.Create(*f, userId.(string))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func MealRetrieve(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	ms := new(services.MealService)
	meal, err := ms.Retrieve(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(meal)
}

func MealUpdate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	f := new(forms.MealForm)
	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	err = ms.Update(*f, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": true})
}

func MealAddIngredient(ctx *fiber.Ctx) error {
	f := new(forms.MealIngredientAddForm)

	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	id, err := ms.AddIngredient(*f)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func MealListIngredients(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	ms := new(services.MealService)
	i, err := ms.ListIngredients(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(i)
}

func MealIngredientUpdate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	f := new(forms.MealIngredientUpdateForm)
	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MealService)
	err = ms.UpdateIngredient(id, *f)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": true})
}

func MealIngredientDelete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	ms := new(services.MealService)
	err := ms.DeleteIngredient(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": true})
}
