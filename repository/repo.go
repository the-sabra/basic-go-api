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

type Storage struct {
	db *gorm.DB
	*GormUserRepo
	*GormBookRepo
}

func NewStorage(dbName string) (*Storage, error) {
	db, err := connectDatabase(dbName)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db:           db,
		GormUserRepo: NewGormUserRepo(db),
		GormBookRepo: NewGormBookRepo(db),
	}, nil
}

// Migrate applies model migrations to the database.
func (s *Storage) Migrate() error {
	var err error

	err = s.db.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}

	err = s.db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
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

// connectDatabase established a new gorm sqlite connection.
func connectDatabase(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
