package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("Requests.go: ")
	log.SetFlags(5)
	log.Println("Invoked")

	invokeServer()
}

// invokeServer starts the HTTP server and sets up the routes and middleware.
func invokeServer() {
	// Middleware Invoke
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", GetTodosHandler)
	mux.HandleFunc("POST /todos", CreateTodoHandler)

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", SetHeader(LogRequest(mux))); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// GetTodosHandler fetches a list of todos from JSON todos.json
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := GetAllTodos()
	HandleResponse(w, todos)
}

// CreateTodoHandler handles the creation of a new todo item. It reads the request body, creates a new todo, and sends the response back to the client.
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var createTodoRequest CreateTodoRequest
	// if err := json.NewDecoder(r.Body).Decode(&createTodoRequest); err != nil {
	// 	HandleErrorResponse(w, err, http.StatusBadRequest)
	// 	return
	// }
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields in the JSON request body
	if err := decoder.Decode(&createTodoRequest); err != nil {
		HandleErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := createTodoRequest.Validate(); err != nil {
		HandleErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	newTodo := CreateTodo(createTodoRequest.Title, createTodoRequest.Completed)
	responseBody := map[string]interface{}{
		"message": "Todo created successfully",
		"todo":    newTodo,
	}

	HandleResponse(w, responseBody)
}

// GetRequest performs a GET request to the specified URL
// func GetRequest() {
// 	log.Println("GetRequest Example")

// 	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err != nil {
// 		log.Fatal("Error encountered: ", err)
// 	}

// 	defer response.Body.Close()
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal("Error reading response body: ", err)
// 	}

// 	log.Println("Response Body:", string(body))
// }
