package controllers

import (
	"net/http"
	"strconv"

	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
)

type LabController struct {
	store *store.MySQLStore
}

func NewLabController(store *store.MySQLStore) *LabController {
	return &LabController{store: store}
}

func (lc *LabController) GetLabs(c *fiber.Ctx) error {
	universityID, err := strconv.ParseInt(c.Query("university_id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid university_id",
		})
	}

	labs, err := lc.store.GetLabsByUniversity(universityID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch labs",
		})
	}

	return c.JSON(fiber.Map{
		"data": labs,
	})
}

func (lc *LabController) GetLab(c *fiber.Ctx) error {
	universityID, err := strconv.ParseInt(c.Query("university_id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid university_id",
		})
	}

	slug := c.Params("slug")

	lab, err := lc.store.GetLabBySlug(universityID, slug)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Lab not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": lab,
	})
}
