package repository

import (
	"context"

	"github.com/shravan/workday/models"
)

type UserRepository interface {
	GetUsers(context.Context) ([]models.User, error)
	GetByID(context.Context, int64) (models.User, error)
	UpdateUser(context.Context, int64, map[string]interface{}) error
	SaveUser(context.Context, *models.User) error
	DeleteUser(context.Context, int64) error
}
