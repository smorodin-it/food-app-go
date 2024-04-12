package handlers

import (
	"errors"
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

func InventoryList(is services.InventoryService) fiber.Handler {
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

func InventoryCreate(is services.InventoryService, as services.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		f := new(forms.InventoryForm)
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
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseAdd(ctx, *id)
	}
}

func InventoryRetrieve(is services.InventoryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if id == "" {
			return utils.GetResponseError(ctx, fiber.StatusBadRequest, errors.New("bad id"))
		}

		inventory, err := is.Retrieve(id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return ctx.Status(fiber.StatusOK).JSON(inventory)
	}
}

func InventoryUpdate(is services.InventoryService) fiber.Handler {
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

		err = is.Update(f, id)
		if err != nil {
			return utils.GetResponseError(ctx, fiber.StatusInternalServerError, err)
		}

		return utils.GetResponseStatus(ctx, true)
	}
}
