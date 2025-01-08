package handlers

import (
	"errors"
	"food-backend/src/forms"
	_ "food-backend/src/responses"
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

// List is a function to get list of all inventory
// @Summary Get inventory list
// @Description Get inventory list
// @Tags Inventory
// @Security ApiKeyAuth
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ []domains.Inventory ]
// @Router /inventory [get]
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

		return utils.GetResponseData(ctx, list)
	}
}

// Create is a function to create new inventory
// @Summary Create new inventory
// @Description Create new inventory
// @Tags Inventory
// @Security ApiKeyAuth
// @Param request body forms.InventoryForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseApi[ responses.ResponseAdd ]
// @Router /inventory [post]
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

// Retrieve is a function to get inventory by id
// @Summary Retrieve inventory by id
// @Description Retrieve inventory by id
// @Tags Inventory
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ domains.Inventory ]
// @Router /inventory/{id} [get]
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

		return utils.GetResponseData(ctx, inventory)
	}
}

// Update is a function to update inventory by id
// @Summary Update inventory
// @Description Update inventory
// @Tags Inventory
// @Security ApiKeyAuth
// @Param request path string true "id"
// @Param request body forms.InventoryForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseApi[ responses.ResponseStatus ]
// @Router /inventory/{id} [put]
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
