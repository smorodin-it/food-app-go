package handlers

import (
	"food-backend/src/forms"
	"food-backend/src/services"
	"github.com/gofiber/fiber/v2"
)

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

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": true})
}