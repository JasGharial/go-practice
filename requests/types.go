package main

// Todo represents a single todo item with an ID, title, and completion status.
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Todos represents a collection of todo items.
type Todos struct {
	Todos []Todo `json:"todos"`
}

// CreateTodoRequest represents the request payload for creating a new todo item, including the title and an optional completion status.
type CreateTodoRequest struct {
	Title     string `json:"title"`
	Completed *bool  `json:"completed"`
}

// UpdateTodoRequest represents the request payload for updating a todo item, including the title and an optional completion status.
type UpdateTodoRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed *bool  `json:"completed"`
}

// DeleteTodoRequest represent the request payload for deleting a todo item.
type DeleteTodoRequest struct {
	ID int `json:"id"`
}
