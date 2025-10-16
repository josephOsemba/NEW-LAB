package api

import (
	"github.com/josephOsemba/go-backend/internal/app/controllers"
	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, store *store.MySQLStore) {
	universityCtrl := controllers.NewUniversityController(store)
	labCtrl := controllers.NewLabController(store)
	experimentCtrl := controllers.NewExperimentController(store)

	api := app.Group("/api/v1")

	api.Get("/universities", universityCtrl.GetUniversities)
	api.Get("/universities/:slug", universityCtrl.GetUniversity)

	labs := api.Group("/labs")
	labs.Get("/", labCtrl.GetLabs)
	labs.Get("/:slug", labCtrl.GetLab)

	experiments := api.Group("/experiments")
	experiments.Get("/", experimentCtrl.GetExperiments)
	experiments.Get("/:slug", experimentCtrl.GetExperiment)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Virtual Lab API is running",
		})
	})
}
