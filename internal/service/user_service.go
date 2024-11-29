package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/models"
)

func (s *service) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if !user.Active {
		return nil, errors.New("user account is inactive")
	}
	return user, nil
}

func (s *service) GetUsers(filters map[string]interface{}) ([]models.User, error) {
	// Add active filter by default unless specified
	if _, ok := filters["active"]; !ok {
		filters["active"] = true
	}
	return s.repo.GetUsers(filters)
}

func (s *service) UpdateUser(id uuid.UUID, updates map[string]interface{}) error {
	// Get existing user
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	// Update user fields
	for key, value := range updates {
		switch key {
		case "first_name":
			user.FirstName = value.(string)
		case "last_name":
			user.LastName = value.(string)
		case "email":
			user.Email = value.(string)
		case "address":
			user.Address = value.(string)
		case "active":
			user.Active = value.(bool)
		}
	}

	return s.repo.UpdateUser(id, user)
}
