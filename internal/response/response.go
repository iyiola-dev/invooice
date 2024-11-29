package response

import "github.com/iyiola-dev/numeris/internal/models"

type LoginResponse struct {
	User  *models.User `json:"user"`
	Token string      `json:"token"`
}