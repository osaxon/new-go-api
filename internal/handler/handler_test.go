package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"new-go-api/internal/database"
	mocks "new-go-api/internal/mocks/service"
	"os"
	"strings"
	"testing"
	"time"
)

func generateTestToken(user database.User) (string, error) {
	// You can mock this if needed, or use a test version
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("test-secret"))
}

func TestHandler_Register(t *testing.T) {
	mockUser := database.User{
		ID:       1,
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}
	mockService := mocks.NewUserServiceMock(t)
	mockService.On("Register", mock.Anything, mock.Anything).Return(mockUser, nil)

	e := echo.New()
	h := NewHandler(mockService)
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"email":"test@example.com", "password":"password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.RegisterUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp database.User
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp)) {
			// Check that the correct data is returned in the response
			assert.Equal(t, mockUser.Email, resp.Email)
			assert.Equal(t, mockUser.Name, resp.Name)
		}
	}

	// Ensure that the mock expectations were met
	mockService.AssertExpectations(t)
}

func TestHandler_Login(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		return
	}
	testPw := os.Getenv("TEST_PW")
	mockUser := database.CreateUserParams{
		Email:    "test@example.com",
		Password: testPw,
		Name:     "Test User",
	}
	mockService := mocks.NewUserServiceMock(t)
	mockService.On("Login", mock.Anything, strings.NewReader(`{"email":"test@example.com", "password":"password"}`)).Return(mockUser, nil)
	e := echo.New()
	h := NewHandler(mockService)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com", "password":"password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the Login handler
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token") // Check if a token is returned
	}

	// Ensure that the mock expectations were met
	mockService.AssertExpectations(t)
}
