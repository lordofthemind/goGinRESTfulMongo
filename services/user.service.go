package services

import "manishh.me/ginAPIs/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUser() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
