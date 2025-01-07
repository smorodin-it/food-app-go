package handlers

import (
	"errors"
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

type InventoryHandler interface {
	List() fiber.Handler
	Create() fiber.Handler
	Retrieve() fiber.Handler
	Update() fiber.Handler
}

type inventoryHandler struct {
	as services.AuthService
	is services.InventoryService
}

func (h inventoryHandler) List() fiber.Handler {
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

		return ctx.Status(fiber.StatusOK).JSON(list)
	}
}

func (h inventoryHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.InventoryForm)
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
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

func (h inventoryHandler) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		inventory, err := h.is.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(inventory)
	}
}

func (h inventoryHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		f := new(forms.InventoryForm)
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

func NewInventoryHandler(as services.AuthService, is services.InventoryService) InventoryHandler {
	return &inventoryHandler{
		as: as,
		is: is,
	}
}
