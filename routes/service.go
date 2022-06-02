package routes

import (
	"context"

	"github.com/shravan/workday/models"
	"github.com/shravan/workday/repository"
	"github.com/shravan/workday/utils"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	return us.repository.GetUsers(ctx)
}
func (us *userService) SaveUser(ctx context.Context, user models.UserRequest) error {
	newUser := models.User{
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		Password:     utils.GenerateHash(user.Password),
	}
	return us.repository.SaveUser(ctx, &newUser)
}
func (us *userService) GetByID(ctx context.Context, id int64) (res models.User, err error) {

	res, err = us.repository.GetByID(ctx, id)
	return
}
func (us *userService) UpdateUser(ctx context.Context, id int64, updatedFileds map[string]interface{}) (user models.User, err error) {

	user, err = us.repository.GetByID(ctx, id)
	if err != nil {
		return user, err
	}
	if user == (models.User{}) {
		return user, err
	}
	return user, us.repository.UpdateUser(ctx, id, updatedFileds)
}
func (us *userService) DeleteUser(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	user, err := us.repository.GetByID(ctx, id)
	if err != nil {
		return user, err
	}
	if user == (models.User{}) {
		return user, err
	}
	return user, us.repository.DeleteUser(ctx, id)
}
