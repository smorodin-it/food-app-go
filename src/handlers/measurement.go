package handlers

import (
	"food-backend/src/forms"
	_ "food-backend/src/responses"
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

// ListByUserId is a function to get list of all user measurements
// @Summary Get measurements list by user id
// @Description Get measurements list by user id
// @Tags Measurement
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []domains.Measurement ]
// @Router /measurement [get]
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

		return utils.GetResponseData(ctx, measurements)
	}
}

// Create is a function to create new measurement
// @Summary Create new measurement
// @Description Create new measurement
// @Tags Measurement
// @Security ApiKeyAuth
// @Param request query forms.MeasurementCreateForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /measurement [post]
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

// Retrieve is a function to get measurement by id
// @Summary Get measurement by id
// @Description Get measurement by id
// @Tags Measurement
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ domains.Measurement ]
// @Router /measurement/{id} [get]
func (h measurementHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		measurement, err := h.ms.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseData(ctx, measurement)
	}
}

// Update is a function to update measurement by id
// @Summary Update measurement
// @Description Update measurement
// @Tags Measurement
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Param request body forms.MeasurementUpdateForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /measurement/{id} [put]
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

// Delete is a function to delete measurement by id
// @Summary Delete measurement
// @Description Delete measurement
// @Tags Measurement
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /measurement/{id} [delete]
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
