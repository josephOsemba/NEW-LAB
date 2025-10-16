package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
}

// TenantMiddleware would extract tenant context from JWT or headers
func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// For now, we'll get tenant from query param
		// In production, this would come from JWT claims
		tenant := c.Query("university_id")
		if tenant == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "university_id required",
			})
		}

		c.Locals("university_id", tenant)
		return c.Next()
	}
}
