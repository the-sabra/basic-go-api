package handlers

import (
	"firstApi/dto"
	"firstApi/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)
type BookHandler struct{
	Repo repository.BookRepository
} 	

func NewBookHandler(repo repository.BookRepository) *BookHandler{
	return &BookHandler{Repo: repo};
}

func (h * BookHandler)CreateBook(c echo.Context) error{
	//get request data from context
	userId := c.Get("user_id")

	//get request data from context
	book := new(dto.Book)

	book.UserId = userId.(uint)
	  
	if err := c.Bind(book); err != nil {
		log.Error("BIND ERROR",err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	} 

	//validate request data
	if err := c.Validate(book); err != nil {
			log.Error("VALIDATE ERROR",err.Error())
			return err
	}

	err:= h.Repo.CreateBook(book)

	if err != nil {
		log.Error("Failed to create book",err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create book")
	}

	return c.JSON(http.StatusCreated, book);
}


func(h *BookHandler)GetBooks(c echo.Context) error{
	books, err := h.Repo.GetAllBook()

	if err != nil {
		log.Error("Failed to get books", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get books")
	}

	return c.JSON(http.StatusOK, books);
}

func (h *BookHandler)GetBook(c echo.Context) error{
	id := c.Param("id")

	//convert id to int
	i,err := strconv.Atoi(id); 
	
	if err != nil{
		log.Error("Failed to convert id to int")
		return c.JSON(http.StatusBadRequest, "Invalid id , id must be number")
	}
	book, err := h.Repo.GetBookById(i)

	if err != nil {
		log.Info("Failed to get book ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get book")
	}

	return c.JSON(http.StatusOK, book);
}


func (h *BookHandler)UpdateBook(c echo.Context) error{
	//get request data from context
	IdParam := c.Param("id")
	user_id := c.Get("user_id").(uint)
	role := c.Get("role").(string)

	//convert id to int"
	bookId , err := strconv.Atoi(IdParam);

	if err != nil{
		log.Error("Failed to convert id to int")
		return c.JSON(http.StatusBadRequest, "Invalid id, id must be number")
	}

	book := new(dto.UpdateBook)

	
	if err := c.Bind(book); err != nil {
		log.Error("BIND ERROR", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	exist ,err := h.Repo.GetBookById(bookId)
	
	if(err != nil){
		log.Error("Failed to get book", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get book")
	}

	if(exist.UserID != user_id && role != "admin"){
		log.Error("You are not authorized to update this book")
		return c.JSON(http.StatusUnauthorized, "You are not authorized to update this book")
	}

	//validate request data
	if err := c.Validate(book); err != nil {
			log.Error("VALIDATE ERROR", err.Error())
			return err
	}

	updatedBook, err := h.Repo.UpdateBook(bookId , book);

	if err != nil {
		log.Error("Failed to update book", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update book")
	}

	return c.JSON(http.StatusOK, updatedBook);
}


func (h * BookHandler)DeleteBook(c echo.Context) error{
	id := c.Param("id")

	//convert id to int
	bookId,err := strconv.Atoi(id);

	if err != nil{
		log.Error("Failed to convert id to int")
		return c.JSON(http.StatusBadRequest, "Invalid id, id must be number")
	}

	err = h.Repo.DeleteBook(bookId)

	if err != nil {
		log.Error("Failed to delete book", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to delete book")
	}
 
	return c.JSON(http.StatusAccepted,map[string]string{"message":"Book Deleted"})
}