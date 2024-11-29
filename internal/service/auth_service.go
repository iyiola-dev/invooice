package service

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/models"
	"github.com/iyiola-dev/numeris/internal/response"
	"github.com/iyiola-dev/numeris/internal/util"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(input inputs.RegisterInput) (*models.User, error) {
	// Check if user exists
	existingUsers, err := s.repo.GetUsers(map[string]interface{}{
		"email": input.Email,
	})
	if err != nil {
		return nil, err
	}
	if len(existingUsers) > 0 {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		ID:        uuid.New(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Address:   input.Address,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Login(input inputs.LoginInput) (*response.LoginResponse, error) {
	// Get user by email
	users, err := s.repo.GetUsers(map[string]interface{}{
		"email": input.Email,
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("invalid email or password")
	}

	user := &users[0]

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.Active {
		return nil, errors.New("account is inactive")
	}

	// Generate JWT token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("error generating authentication token")
	}

	// Create activity log
	activityLog := &models.ActivityLog{
		UserID:    user.ID,
		Action:    "LOGIN",
		Timestamp: time.Now(),
	}
	
	err = s.repo.CreateActivityLog(activityLog)
	if err != nil {
		// Log the error but don't fail the login
		log.Printf("Failed to create activity log: %v", err)
	}

	return &response.LoginResponse{
		User:  user,
		Token: token,
	}, nil
}
