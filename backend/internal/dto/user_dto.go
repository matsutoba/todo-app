package dto

type CreateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Passowrd string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Passowrd string `json:"password" binding:"required"`
}
