package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/mocks"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	input := inputs.RegisterInput{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
		Password:  "password123",
	}

	// Set up expectations
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	// Execute
	user, err := svc.Register(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.FirstName, user.FirstName)
	assert.Equal(t, input.Email, user.Email)

	// Verify password was hashed
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestRegister_CreateUserError(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	input := inputs.RegisterInput{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@example.com",
		Password:  "password123",
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(errors.New("duplicate email"))

	user, err := svc.Register(input)

	assert.Error(t, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	existingUser := &models.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	input := inputs.LoginInput{
		Email:    "test@example.com",
		Password: password,
	}

	mockRepo.On("GetUsers", map[string]interface{}{"email": input.Email}).Return([]models.User{*existingUser}, nil)

	resp, err := svc.Login(input)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)
	assert.Equal(t, existingUser.ID, resp.User.ID)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	input := inputs.LoginInput{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	mockRepo.On("GetUsers", map[string]interface{}{"email": input.Email}).Return([]models.User{}, nil)

	resp, err := svc.Login(input)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "invalid credentials", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.Repository)
	svc := service.NewService(mockRepo)

	input := inputs.LoginInput{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	mockRepo.On("GetUsers", map[string]interface{}{"email": input.Email}).Return([]models.User{}, nil)

	resp, err := svc.Login(input)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "invalid credentials", err.Error())
	mockRepo.AssertExpectations(t)
}
