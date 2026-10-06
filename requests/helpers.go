package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"
)

// HandleResponse writes the provided data as a JSON response to the http.ResponseWriter.
func HandleResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding response to JSON"))
	}
}

// HandleErrorResponse writes an error message and status code to the http.ResponseWriter.
func HandleErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	log.Println("Error:", err)
	w.WriteHeader(statusCode)

	response := map[string]string{"error": err.Error()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding error response to JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding error response to JSON"))
	}
}

// ReadJSON reads the todos.json file and returns its content as a map[string]interface{}.
func ReadJSON() map[string]interface{} {
	jsonFile, er := os.Open("todos.json")
	if er != nil {
		panic(er)
	}
	defer jsonFile.Close()

	bytevalue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytevalue, &result)

	return result
}

// WriteJSON writes the provided data to the todos.json file in JSON format.
func WriteJSON(data map[string]interface{}) {
	jsonFile, err := os.Create("todos.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

// Receiver Functions for HTTP requests and responses.
// Validate validates the CreateTodoRequest struct to ensure that the title is not empty and does not exceed 100 characters.
func (r CreateTodoRequest) Validate() error {
	if strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("title is required")
	}
	if len(r.Title) > 100 {
		return fmt.Errorf("title must not exceed more than 100 characters")
	}
	return nil
}