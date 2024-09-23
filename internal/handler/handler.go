package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"new-go-api/internal/database"
	"new-go-api/internal/service"
)

type Handler struct {
	UserService service.UserService
}

func NewHandler(us service.UserService) *Handler {
	return &Handler{
		UserService: us,
	}
}

func (h *Handler) HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
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

func (h *Handler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	token, err := h.UserService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) RegisterUser(c echo.Context) error {
	// Hash the user's password
	ctx := c.Request().Context()
	var req RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	userParams := database.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	// Create the user in the database
	user, err := h.UserService.Register(ctx, userParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
