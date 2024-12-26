package main

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kamran0812/notes-backend/conf"
	"github.com/kamran0812/notes-backend/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	//User routes
	api.Post("/signup", handler.Signup)
	api.Post("/login", handler.Login)

	//Note routes
	note := api.Group("/note", AuthMiddleware)
	// note.Use(AuthMiddleware)
	note.Post("/", handler.CreateNotes)
	note.Get("/search", handler.SearchNotes)
	note.Get("/", handler.GetNotes)
	note.Get("/:noteId", handler.GetNote)
	note.Put("/:noteId", handler.UpdateNote)
	note.Delete("/:noteId", handler.DeleteNote)
}

var jwtKey = []byte(conf.GetConfig("AUTH_SECRET"))

func AuthMiddleware(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id := (claims["user_id"])
		c.Locals("user_id", user_id)
		return c.Next()
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
}
