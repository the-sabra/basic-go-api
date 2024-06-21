package repository

import (
	"firstApi/dto"
	"firstApi/models"
)

type UserRepository interface{
	CreateUser(user *dto.User) (*models.User,error)
	GetAllUsers() ([]models.User, error) 
	GetUserById(id uint) (models.User, error)
	UpdateUser(id uint,user dto.User) (*models. User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (models.User,error)
}