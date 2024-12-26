package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/kamran0812/notes-backend/db"
)

func main() {
	app := fiber.New()
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	db.ConnectDB()
	SetupRoutes(app)
	app.Listen(":3000")
}
