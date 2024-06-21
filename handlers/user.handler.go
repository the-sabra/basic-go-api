package handlers

import (
	"firstApi/dto"
	"firstApi/repository"
	"firstApi/util"

	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo};
}

func (h *UserHandler)SignUp(c echo.Context) error {
	user :=new(dto.User)

	if err:= c.Bind(user); err!=nil{
		log.Error("BIND ERROR",err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	} 

	
	if err := c.Validate(user); err!=nil{
		log.Error("VALIDATE ERROR", err.Error())
		return err;
	}

	_,err := h.Repo.GetUserByEmail(user.Email)

	if(err ==nil){
		log.Error("USER ALREADY EXIST")
		return c.JSON(http.StatusBadRequest, "User already exist")
	}   
	 
	hash,err:= util.HashPassword(user.Password);
		if err!=nil{
			log.Error("HASH ERROR", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}
		
	user.Password = hash;
	
	createdUser,err := h.Repo.CreateUser(user);

	if err != nil{
		log.Error("CREATE USER ERROR", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdUser)
}


func(h *UserHandler)Login(c echo.Context) error{
	user := new(dto.LoginUser)

	if err:= c.Bind(user); err!=nil{
		log.Error("BIND ERROR", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(user); err!=nil{
		log.Error("VALIDATE ERROR", err.Error())
		return err;
	}

	
	exist,err := h.Repo.GetUserByEmail(user.Email)
	
	if(err != nil){
		log.Error("USER NOT FOUND")
		return c.JSON(http.StatusBadRequest, "Email or password is wrong")
	} 
	 
	if !util.ComparePassword(user.Password, exist.Password){
		log.Error("PASSWORD MISMATCH")
		return c.JSON(http.StatusUnauthorized, "Email or password is wrong")
	}


	claims := util.JwtCustomClaims{
		UserId: exist.ID,
		Role: exist.Role,
	} 

	token:= util.GenerateJWT(claims);
	if token == ""{
		log.Error("JWT TOKEN ERROR")
		return c.JSON(http.StatusInternalServerError, "JWT token error")
	}

	return c.JSON(http.StatusOK, map[string]any{"token":token,"user":exist})
} 

func(h *UserHandler)GetAllUsers(c echo.Context) error{
	users, err:= h.Repo.GetAllUsers();

	if err!=nil{
		log.Error("GET ALL USERS ERROR", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func(h *UserHandler)GetUser(c echo.Context) error {
	claims := c.Get("claims")
 
	return c.JSON(http.StatusOK, claims)
}   