package repository

import (
	"firstApi/dto"
	"firstApi/models"

	"gorm.io/gorm"
)

type GormBookRepo struct{
	DB *gorm.DB
}



func NewGormBookRepo(DB *gorm.DB) *GormBookRepo{
	return &GormBookRepo{DB: DB}
}



func (r *GormBookRepo) CreateBook(book *dto.Book) error {
	return r.DB.Create(book).Error
}

func (r *GormBookRepo) GetAllBook() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *GormBookRepo) GetBookById(id int) (models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	return book, err
}

func (r *GormBookRepo) UpdateBook(id int, updatedBook *dto.UpdateBook) (*models.Book, error) {
	// Parse and validate the ID
	book, err := r.GetBookById(id)
	if err != nil {
		return nil, err
	}
	// Update the book fields if they are provided in the request
	if updatedBook.Title != "" {
		book.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}

	err = r.DB.Save(book).Error

	return &book, err

} 
func (r *GormBookRepo) DeleteBook(id int) error {
	return r.DB.Delete(&models.Book{}, id).Error
}