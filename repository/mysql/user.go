package mysql

import (
	"context"

	"github.com/shravan/workday/models"
	"github.com/shravan/workday/repository"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserReposity(conn *gorm.DB) repository.UserRepository {
	return &mysqlUserRepository{conn}
}

func (m *mysqlUserRepository) GetUsers(ctx context.Context) (res []models.User, err error) {

	query := m.Conn.Find(&res)
	err = query.Error
	return
}

func (m *mysqlUserRepository) SaveUser(ctx context.Context, user *models.User) (err error) {
	err = m.Conn.Create(&user).Error
	return
}
func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (res models.User, err error) {
	query := m.Conn.Find(&res, id)
	err = query.Error

	return
}



func (m *mysqlUserRepository) DeleteUser(ctx context.Context, id int64) (err error) {
	err = m.Conn.Delete(models.User{}, id).Error
	return
}
func (m *mysqlUserRepository) UpdateUser(ctx context.Context, id int64, updatedFields map[string]interface{}) (err error) {
	err = m.Conn.Model(models.User{}).Where("id=?", id).Updates(updatedFields).Error
	return
}