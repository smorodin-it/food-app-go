package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"food-backend/src/utils"
	"github.com/gofiber/fiber/v2"
)

// IngredientList is a function to get list of all ingredients
// @Summary Get ingredients list
// @Description Get ingredients list
// @Tags Ingredient
// @Param request query forms.PaginationQuery true "pagination"
// @Produce json
// @Success 200 {object} []domains.Ingredient
// @Router /ingredient [get]
func IngredientList(ctx *fiber.Ctx) error {
	q := new(forms.PaginationQuery)

	err := ctx.QueryParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	is := new(services.IngredientService)

	list, err := is.List(q.Page, q.PerPage)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

// IngredientCreate is a function to create new ingredient
// @Summary Create new ingredient
// @Description Create new ingredient
// @Tags Ingredient
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 201 {object} responses.ResponseAdd
// @Router /ingredient [post]
func IngredientCreate(ctx *fiber.Ctx) error {
	f := new(forms.IngredientForm)
	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	as := new(services.AuthService)
	userId, err := as.GetUserIdFromToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	is := new(services.IngredientService)
	id, err := is.Create(f, userId.(string))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.GetResponseAdd(ctx, *id)
}

// IngredientRetrieve is a function to get ingredient by id
// @Summary Retrieve ingredient by id
// @Description Retrieve ingredient by id
// @Tags Ingredient
// @Param request path string true "id"
// @Produce json
// @Success 200 {object} domains.Ingredient
// @Router /ingredient/${id} [get]
func IngredientRetrieve(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	is := new(services.IngredientService)

	ingredient, err := is.Retrieve(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ingredient)
}

// IngredientUpdate is a function to update ingredient by id
// @Summary Update ingredient
// @Description Update ingredient
// @Tags Ingredient
// @Param request path string true "id"
// @Param request body forms.IngredientForm true "body"
// @Produce json
// @Success 200 {object} responses.ResponseStatus
// @Router /ingredient/${id} [put]
func IngredientUpdate(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad id"})
	}

	f := new(forms.IngredientForm)

	err := ctx.BodyParser(f)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	is := new(services.IngredientService)

	err = is.Update(f, id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return utils.GetResponseStatus(ctx, true)
}
