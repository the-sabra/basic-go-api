package repository

import (
	"firstApi/dto"
	"firstApi/models"

	"gorm.io/gorm"
)

type GormUserRepo struct{ 
	DB *gorm.DB
}



func NewGormUserRepo(DB *gorm.DB) *GormUserRepo{
	return &GormUserRepo{DB: DB}
}



func (r *GormUserRepo) CreateUser(dtoUser *dto.User) (*models.User, error) {
	// Convert dto.User to models.User
	modelUser := models.User{
		Name: dtoUser.Name,
		Email:    dtoUser.Email,
		Password: dtoUser.Password,
	}

	if err := r.DB.Create(&modelUser).Error; err != nil {
		return nil, err
	}

	// Do not return the password
	modelUser.Password = ""

	return &modelUser, nil
} 

 func(r *GormUserRepo) GetAllUsers() ([]models.User, error){
	var users []models.User
	err:= r.DB.Find(&users).Error
	return users, err
 }

 func(r *GormUserRepo) GetUserById(id uint) (models.User, error){
	var user models.User
	err:= r.DB.Where("id = ?", id).First(&user).Error
	return user, err
 }	

 func(r *GormUserRepo) UpdateUser(id uint,updatedUser dto.User) (*models.User,error){
	user, err := r.GetUserById(id)
	if err != nil {
		return nil,err
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}

	if(updatedUser.Email != ""){
		user.Email = updatedUser.Email
	}
	
	// Update other fields as needed

	err = r.DB.Save(&user).Error

	// Handle any errors that may occur during the update process
	return &user,err
	
}

 func(r *GormUserRepo) DeleteUser(id uint) (error){
	err:= r.DB.Where("id = ?", id).Delete(&models.User{}).Error
	return err
 }

 func(r*GormUserRepo)GetUserByEmail(email string) (models.User,error){
	var user models.User
	 err :=r.DB.Where("email = ?", email).First(&user).Error
	return user ,err  
 } 