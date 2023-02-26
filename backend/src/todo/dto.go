package todo

type CreateTodoDto struct {
	Title string `json:"title" binding:"required"`
}

type UpdateTodoDto struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
