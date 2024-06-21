package dto
 
type Book struct{
	Title string `json:"title" validate:"required,min=2,max=50"`
	Author string `json:"author" validate:"required,min=2,max=50"`
	UserId uint `json:"user_id" validate:"omitempty,number"`
}
 
type UpdateBook struct {
	Title string `json:"title" validate:"omitempty,min=2,max=10"`
	Author string `json:"author" validate:"omitempty,min=2,max=20"`
	UserId uint `json:"user_id" validate:"omitempty,number"`
}   