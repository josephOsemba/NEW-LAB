package api

import (
	"github.com/josephOsemba/go-backend/internal/app/controllers"
	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, store *store.MySQLStore) {
	// Initialize controllers
	universityCtrl := controllers.NewUniversityController(store)
	labCtrl := controllers.NewLabController(store)
	experimentCtrl := controllers.NewExperimentController(store)

	// API v1 routes
	api := app.Group("/api/v1")

	// Universities
	api.Get("/universities", universityCtrl.GetUniversities)
	api.Get("/universities/:slug", universityCtrl.GetUniversity)

	// Labs
	labs := api.Group("/labs")
	labs.Get("/", labCtrl.GetLabs)
	labs.Get("/:slug", labCtrl.GetLab)

	// Experiments
	experiments := api.Group("/experiments")
	experiments.Get("/", experimentCtrl.GetExperiments)
	experiments.Get("/:slug", experimentCtrl.GetExperiment)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Virtual Lab API is running",
		})
	})
}
