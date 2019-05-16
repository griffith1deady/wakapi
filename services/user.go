package services

import (
	"github.com/jinzhu/gorm"
	"github.com/n1try/wakapi/models"
)

const TableUser = "user"

type UserService struct {
	Db *gorm.DB
}

func (srv *UserService) GetUserById(userId string) (*models.User, error) {
	u := &models.User{}
	if err := srv.Db.Where(&models.User{ID: userId}).First(u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (srv *UserService) GetUserByKey(key string) (*models.User, error) {
	u := &models.User{}
	if err := srv.Db.Where(&models.User{ApiKey: key}).First(u).Error; err != nil {
		return u, err
	}
	return u, nil
}