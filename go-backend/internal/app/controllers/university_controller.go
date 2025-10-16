package controllers

import (
	"net/http"

	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
)

type UniversityController struct {
	store *store.MySQLStore
}

func NewUniversityController(store *store.MySQLStore) *UniversityController {
	return &UniversityController{store: store}
}

func (uc *UniversityController) GetUniversities(c *fiber.Ctx) error {
	universities, err := uc.store.GetUniversities()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch universities",
		})
	}

	return c.JSON(fiber.Map{
		"data": universities,
	})
}

func (uc *UniversityController) GetUniversity(c *fiber.Ctx) error {
	slug := c.Params("slug")

	university, err := uc.store.GetUniversityBySlug(slug)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "University not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": university,
	})
}
