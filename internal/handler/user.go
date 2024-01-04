package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kamran0812/notes-backend/conf"
	"github.com/kamran0812/notes-backend/db"
	"github.com/kamran0812/notes-backend/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	db := db.Conn
	var newUser model.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	newUser.Password, _ = hashPassword(newUser.Password)
	if err := db.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	db := db.Conn

	var creds model.User
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var user model.User
	if err := db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !ValidatePassword(creds.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, _ := generateToken(user.ID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func generateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})

	secretKey := []byte(conf.GetConfig("AUTH_SECRET"))
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ValidatePassword(inputPassword string, storedHashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(inputPassword))
	return err == nil
}
