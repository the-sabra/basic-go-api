package repository

import (
	"firstApi/dto"
	"firstApi/models"
)

type BookRepository interface {
	CreateBook(book *dto.Book) (error)
	GetAllBook() ([]models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(id int, book *dto.UpdateBook) (*models.Book, error)
	DeleteBook(id int) error
} 