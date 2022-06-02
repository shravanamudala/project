package routes

import (
	"context"
	"net/http"

	"github.com/shravan/workday/models"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	SaveUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserService interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int64) (models.User, error)
	SaveUser(ctx context.Context, user models.UserRequest) error
	UpdateUser(ctx context.Context, id int64, data map[string]interface{}) (models.User, error)
	DeleteUser(ctx context.Context, id int64) (models.User, error)
}
