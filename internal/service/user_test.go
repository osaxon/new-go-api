package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"new-go-api/internal/database"
	mocks "new-go-api/internal/mocks/database"
	"os"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	// arrange
	err := godotenv.Load()
	if err != nil {
		return
	}
	testPw := os.Getenv("TEST_PW")
	mockUser := database.User{
		ID:       1,
		Email:    "test@example.com",
		Password: testPw,
		Name:     "Test User",
	}
	mockQueries := mocks.NewQuerierMock(t)
	mockQueries.On("GetUserByEmail", mock.Anything, "test@example.com").Return(mockUser, nil)
	ctx := context.Background()
	service := NewUserService(mockQueries)

	// act
	user, err := service.Login(ctx, "test@example.com", "password")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
	mockQueries.AssertNumberOfCalls(t, "GetUserByEmail", 1)
	mockQueries.AssertExpectations(t)
}
