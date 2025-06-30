package main

import (
	"encoding/json"
	"log"
	"net/http"

	auto "github.com/dani/go-auto-swagger/go-auto-swagger"
)

// User represents a user in the system.
type User struct {
	ID    int    `json:"id" description:"User ID" example:"1"`
	Name  string `json:"name" description:"User name" example:"John Doe"`
	Email string `json:"email" description:"User email" example:"john@example.com"`
}

// CreateUserRequest represents the request to create a new user.
type CreateUserRequest struct {
	Name  string `json:"name" description:"User name" example:"John Doe"`
	Email string `json:"email" description:"User email" example:"john@example.com"`
}

// CreateUserResponse represents the response after creating a new user.
type CreateUserResponse struct {
	User    User   `json:"user" description:"Created user"`
	Message string `json:"message" description:"Success message" example:"User created successfully"`
}

func main() {
	// Create AutoSwagger instance with net/http adapter
	swagger, adapter := auto.NewWithNetHTTP("User API", "1.0.0")
	
	// Configure API info
	swagger.SetInfo(auto.Info{
		Title:       "User Management API",
		Description: "A simple API for managing users",
		Version:     "1.0.0",
		Contact: auto.Contact{
			Name:  "API Support",
			Email: "support@example.com",
			URL:   "https://example.com/support",
		},
		License: auto.License{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
	})
	
	// Add server info
	swagger.AddServer("http://localhost:8080", "Development server")
	
	// Add tags
	swagger.AddTag("users", "User management operations")
	
	// Register routes with full documentation
	adapter.POST("/users", createUserHandler).
		Summary("Create a new user").
		Description("Creates a new user with the provided information").
		Tags("users").
		Request(&CreateUserRequest{}).
		Response(&CreateUserResponse{})
	
	adapter.GET("/users/{id}", getUserHandler).
		Summary("Get user by ID").
		Description("Retrieves a user by their ID").
		Tags("users").
		Path("id", "User ID", auto.StringSchema).
		Response(&User{})
	
	adapter.GET("/users", listUsersHandler).
		Summary("List all users").
		Description("Retrieves a list of all users").
		Tags("users").
		Query("limit", "Number of users to return", false, auto.IntSchema).
		Query("offset", "Number of users to skip", false, auto.IntSchema).
		Response(&[]User{})
	
	// Serve Swagger documentation
	adapter.ServeDocs("/swagger.json")
	adapter.ServeSwaggerUI("/docs", "/swagger.json")
	
	// Simple hello endpoint
	adapter.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
	}).
		Summary("Hello World").
		Description("Returns a simple greeting").
		Tags("general")

	log.Println("Server starting on :8080")
	log.Println("API Documentation: http://localhost:8080/docs")
	log.Println("Swagger JSON: http://localhost:8080/swagger.json")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// Simulate user creation
	user := User{
		ID:    1,
		Name:  req.Name,
		Email: req.Email,
	}
	
	resp := CreateUserResponse{
		User:    user,
		Message: "User created successfully",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path (simplified)
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
