package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

func MeasurementListByUserId(ctx *fiber.Ctx) error {
	userId := utils.GetUserIDFromToken(ctx)

	q := new(forms.PaginationQuery)
	err := ctx.QueryParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MeasurementService)
	measurements, err := ms.ListByUserId(q.Page, q.PerPage, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(measurements)
}

func MeasurementCreate(ctx *fiber.Ctx) error {
	f := new(forms.MeasurementCreateForm)
	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ms := new(services.MeasurementService)
	id, err := ms.Create(*f)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.GetResponseAdd(ctx, *id)
}
