package dto

type User struct {
	Name     string `json:"name" validate:"required,min=2,max=10"`
	Email    string `json:"email" gorm:"unique" validate:"email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
	Role     string `json:"role"`
}  


type UpdateUser struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=15"`
	Email    string `json:"email" validate:"omitempty,email"`
	Role     string `json:"-"`
}


type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
} 