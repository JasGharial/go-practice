package main

import "fmt"

// GetAllTodos retrieves all todo items from the todos.json file and returns them as a slice of Todo structs.
func GetAllTodos() []Todo {
	var todos []Todo
	jsonData := ReadJSON()

	if todosData, ok := jsonData["todos"].([]interface{}); ok {
		for _, todoData := range todosData {
			if todoMap, ok := todoData.(map[string]interface{}); ok {
				todo := Todo{
					ID:        int(todoMap["id"].(float64)),
					Title:     todoMap["title"].(string),
					Completed: todoMap["completed"].(bool),
				}
				todos = append(todos, todo)
			}
		}
	}

	if todos == nil {
		todos = []Todo{}
	}

	return todos
}

// CreateTodo creates a new todo item with the provided title and completion status, adds it to the existing list of todos, and writes the updated list back to the todos.json file. It returns the newly created Todo item.
func CreateTodo(title string, completed *bool) Todo {
	if completed == nil {
		completed = new(bool)
		*completed = false
	}

	todos := GetAllTodos()
	newID := len(todos) + 1
	newTodo := Todo{
		ID:        newID,
		Title:     title,
		Completed: *completed,
	}

	todos = append(todos, newTodo)

	jsonData := map[string]interface{}{
		"todos": todos,
	}
	WriteJSON(jsonData)

	return newTodo
}

// UpdateTodo updates an existing todo item with the specified ID, changing its title and completion status. It returns the updated Todo item and an error if the todo with the given ID is not found.
func UpdateTodo(id int, title string, completed bool) (Todo, error) {
	todos := GetAllTodos()
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = title
			todos[i].Completed = completed

			jsonData := map[string]interface{}{
				"todos": todos,
			}
			WriteJSON(jsonData)

			return todos[i], nil
		}
	}

	return Todo{}, fmt.Errorf("todo with ID %d not found", id)
}

// DeleteTodo omits a todo with the id passed in the argument thus deleting a todo. It returns the deleted Todo item and an error if the todo with the given ID is not found.
func DeleteTodo(id int) (Todo, error) {
	todos := GetAllTodos()

	for i, todo := range todos {
		if todo.ID == id {
			// Store the todo before removing it.
			deletedTodo := todo

			// Remove the todo at index i.
			todos = append(todos[:i], todos[i+1:]...)

			// Persist the updated slice.
			jsonData := map[string]interface{}{
				"todos": todos,
			}
			WriteJSON(jsonData)

			return deletedTodo, nil
		}
	}

	return Todo{}, fmt.Errorf("todo with ID %d not found", id)
}
