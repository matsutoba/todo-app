package dto

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	DueDate     string `json:"due_date" validate:"datetime=2006-01-02"`
}

type UpdateTodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	DueDate     string `json:"due_date" validate:"datetime=2006-01-02"`
}
