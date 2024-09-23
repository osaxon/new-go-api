package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"new-go-api/internal/database"
	"os"
	"time"
)

type UserService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	Register(ctx context.Context, user database.CreateUserParams) (database.User, error)
}

type userService struct {
	queries database.Querier
}

func NewUserService(q database.Querier) UserService {
	return &userService{queries: q}
}

func (s *userService) Register(ctx context.Context, user database.CreateUserParams) (database.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return database.User{}, err
	}

	// Check if the user already exists
	_, err = s.queries.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return database.User{}, errors.New("user already exists")
	}

	user.Password = string(hashedPassword)

	// Create a new user in the database
	newUser, err := s.queries.CreateUser(ctx, user)
	if err != nil {
		return database.User{}, err
	}

	return newUser, nil
}

func (s *userService) Login(ctx context.Context, email string, password string) (string, error) {
	// Get the user by email
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare the password
	err = s.verifyPassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	t, err := s.generateJWT(user)

	return t, nil
}

func (s *userService) verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *userService) generateJWT(user database.User) (string, error) {
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
