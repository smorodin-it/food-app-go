package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

func MeasurementListByUserId(ms services.MeasurementService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := utils.GetUserIDFromToken(ctx)

		q := new(forms.PaginationQuery)
		err := ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		measurements, err := ms.ListByUserId(q.Page, q.PerPage, userId)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(measurements)
	}
}

func MeasurementCreate(ms services.MeasurementService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MeasurementCreateForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := ms.Create(*f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

func MeasurementRetrieve(ms services.MeasurementService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		measurement, err := ms.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(measurement)
	}
}

func MeasurementUpdate(ms services.MeasurementService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		f := new(forms.MeasurementUpdateForm)
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

func MeasurementDelete(ms services.MeasurementService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		err := ms.Delete(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}

}
