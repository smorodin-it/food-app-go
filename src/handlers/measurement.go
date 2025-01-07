package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

type MeasurementHandler interface {
	ListByUserId() fiber.Handler
	Create() fiber.Handler
	Retrieve() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type measurementHandler struct {
	as services.AuthService
	ms services.MeasurementService
}

func (h measurementHandler) ListByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId, err := h.as.GetUserIdFromToken(ctx)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		q := new(forms.PaginationQuery)
		err = ctx.QueryParser(q)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		measurements, err := h.ms.ListByUserId(q.Page, q.PerPage, userId.(string))
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(measurements)
	}
}

func (h measurementHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.MeasurementCreateForm)
		err := ctx.BodyParser(f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, err)
		}

		id, err := h.ms.Create(*f)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

func (h measurementHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		measurement, err := h.ms.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(measurement)
	}
}

func (h measurementHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		f := new(forms.MeasurementUpdateForm)
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

func (h measurementHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		err := h.ms.Delete(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}

func NewMeasurementHandler(as services.AuthService, ms services.MeasurementService) MeasurementHandler {
	return &measurementHandler{
		as: as,
		ms: ms,
	}
}
