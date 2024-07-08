package repository

import (
	"firstApi/dto"
	"firstApi/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Repository represents a generic storage.
type Repository interface {
	UserRepository
	BookRepository
}

// UserRepository represents the functionality needed to be implemented by
// any user interacting repo.
type UserRepository interface {
	CreateUser(user *dto.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	UpdateUser(id uint, user dto.User) (*models.User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (models.User, error)
}

// BookRepository represents the functionality needed to be implemented by
// any book interacting repo.
type BookRepository interface {
	CreateBook(book *dto.Book) error
	GetAllBook() ([]models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(id int, book *dto.UpdateBook) (*models.Book, error)
	DeleteBook(id int) error
}

var DB *gorm.DB

// ConnectDatabase established a new gorm sqlite connection.
func ConnectDatabase(dbName string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = database

	return DB, nil
}

// Migrate applies model migrations to the database.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})
}
