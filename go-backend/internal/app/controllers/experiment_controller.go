package controllers

import (
	"net/http"
	"strconv"

	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
)

type ExperimentController struct {
	store *store.MySQLStore
}

func NewExperimentController(store *store.MySQLStore) *ExperimentController {
	return &ExperimentController{store: store}
}

func (ec *ExperimentController) GetExperiments(c *fiber.Ctx) error {
	labID, err := strconv.ParseInt(c.Query("lab_id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid lab_id",
		})
	}

	experiments, err := ec.store.GetExperimentsByLab(labID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch experiments",
		})
	}

	return c.JSON(fiber.Map{
		"data": experiments,
	})
}

func (ec *ExperimentController) GetExperiment(c *fiber.Ctx) error {
	labID, err := strconv.ParseInt(c.Query("lab_id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid lab_id",
		})
	}

	slug := c.Params("slug")

	experiment, err := ec.store.GetExperimentBySlug(labID, slug)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Experiment not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": experiment,
	})
}
