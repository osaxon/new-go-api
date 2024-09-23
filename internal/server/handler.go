package server

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"new-go-api/internal/database"
	"os"
	"time"
)

func (app *App) HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func (app *App) TestHandler(c echo.Context) error {
	ctx := context.Background()
	newUser, err := app.Q.CreateUser(ctx, database.CreateUserParams{
		Name:     "Oliver",
		Email:    "oli@gmail.com",
		Password: "google21",
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, newUser)

}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (app *App) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, err := app.Q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	if err := verifyPassword(user.Password, req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	token, err := generateJWT(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (app *App) RegisterUser(c echo.Context) error {
	// Hash the user's password
	ctx := c.Request().Context()
	var req RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}
	hashedPassword, err := hashPassword(req.Password)

	if err != nil {
		return err
	}

	// Create the user in the database
	user, err := app.Q.CreateUser(ctx, database.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	return c.JSON(http.StatusOK, user)
}

func generateJWT(user database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID, // Subject (user ID)
		"email": user.Email,
		"name":  user.Name,
		"iat":   time.Now().Unix(),                     // Issued At
		"exp":   time.Now().Add(24 * time.Hour).Unix(), // Expires in 24 hours
		"jti":   uuid.New().String(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("CRYPTO_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
