package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
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
