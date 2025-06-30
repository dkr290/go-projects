package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

// UpdateUserRequest represents the request to update a user.
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" description:"User name" example:"John Doe"`
	Email string `json:"email,omitempty" description:"User email" example:"john@example.com"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error" description:"Error message"`
	Code  int    `json:"code" description:"Error code"`
}

func main() {
	// Create chi router
	r := chi.NewRouter()
	
	// Add chi middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	
	// Create AutoSwagger with chi adapter
	swagger, adapter := auto.NewWithChi("User API with Chi", "1.0.0", r)
	
	// Configure API info
	swagger.SetInfo(auto.Info{
		Title:       "User Management API with Chi",
		Description: "A RESTful API for managing users built with Chi router and automatic Swagger documentation",
		Version:     "1.0.0",
		Contact: auto.Contact{
			Name:  "Development Team",
			Email: "dev@example.com",
			URL:   "https://example.com",
		},
		License: auto.License{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
	})
	
	// Add servers
	swagger.AddServer("http://localhost:8080", "Development server")
	swagger.AddServer("https://api.example.com", "Production server")
	
	// Add tags
	swagger.AddTag("users", "User management operations")
	swagger.AddTag("health", "Health and monitoring endpoints")
	
	// Health check endpoint
	adapter.GET("/health", healthHandler).
		Summary("Health Check").
		Description("Returns the health status of the API").
		Tags("health").
		Response(map[string]string{"status": "ok"})
	
	// User routes
	r.Route("/api/v1/users", func(r chi.Router) {
		// Create user
		adapter.POST("/api/v1/users", createUserHandler).
			Summary("Create User").
			Description("Creates a new user with the provided information").
			Tags("users").
			Request(&CreateUserRequest{}).
			Response(&CreateUserResponse{}).
			Responses(map[string]auto.Response{
				"201": auto.JSONResponse("User created successfully", &CreateUserResponse{}),
				"400": auto.ErrorResponse("Invalid request data"),
				"409": auto.ErrorResponse("User already exists"),
				"500": auto.ErrorResponse("Internal server error"),
			})
		
		// List users
		adapter.GET("/api/v1/users", listUsersHandler).
			Summary("List Users").
			Description("Retrieves a paginated list of users").
			Tags("users").
			Query("page", "Page number", false, auto.IntSchema).
			Query("limit", "Number of users per page", false, auto.IntSchema).
			Query("search", "Search term for filtering users", false, auto.StringSchema).
			Response(&[]User{}).
			Responses(map[string]auto.Response{
				"200": auto.JSONResponse("List of users", &[]User{}),
				"400": auto.ErrorResponse("Invalid query parameters"),
				"500": auto.ErrorResponse("Internal server error"),
			})
		
		r.Route("/{userID}", func(r chi.Router) {
			// Get user by ID
			adapter.GET("/api/v1/users/{userID}", getUserHandler).
				Summary("Get User").
				Description("Retrieves a specific user by their ID").
				Tags("users").
				Path("userID", "User ID", auto.IntSchema).
				Response(&User{}).
				Responses(map[string]auto.Response{
					"200": auto.JSONResponse("User details", &User{}),
					"404": auto.ErrorResponse("User not found"),
					"500": auto.ErrorResponse("Internal server error"),
				})
			
			// Update user
			adapter.PUT("/api/v1/users/{userID}", updateUserHandler).
				Summary("Update User").
				Description("Updates an existing user's information").
				Tags("users").
				Path("userID", "User ID", auto.IntSchema).
				Request(&UpdateUserRequest{}).
				Response(&User{}).
				Responses(map[string]auto.Response{
					"200": auto.JSONResponse("User updated successfully", &User{}),
					"400": auto.ErrorResponse("Invalid request data"),
					"404": auto.ErrorResponse("User not found"),
					"500": auto.ErrorResponse("Internal server error"),
				})
			
			// Delete user
			adapter.DELETE("/api/v1/users/{userID}", deleteUserHandler).
				Summary("Delete User").
				Description("Deletes a user from the system").
				Tags("users").
				Path("userID", "User ID", auto.IntSchema).
				Responses(map[string]auto.Response{
					"204": {Description: "User deleted successfully"},
					"404": auto.ErrorResponse("User not found"),
					"500": auto.ErrorResponse("Internal server error"),
				})
		})
	})
	
	// Serve documentation
	adapter.ServeDocs("/swagger.json")
	adapter.ServeSwaggerUI("/docs", "/swagger.json")
	
	// Additional documentation endpoints
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs", http.StatusFound)
	})
	
	log.Println("🚀 Server starting on :8080")
	log.Println("📚 API Documentation: http://localhost:8080/docs")
	log.Println("📋 Swagger JSON: http://localhost:8080/swagger.json")
	log.Println("❤️  Health Check: http://localhost:8080/health")
	log.Println("👥 Users API: http://localhost:8080/api/v1/users")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "ok",
		"timestamp": "2024-01-01T00:00:00Z",
		"version":   "1.0.0",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// Validate request
	if req.Name == "" || req.Email == "" {
		sendError(w, "Name and email are required", http.StatusBadRequest)
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

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	
	page := 1
	limit := 10
	
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}
	
	// Simulate user list (in real app, this would query database)
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		{ID: 3, Name: "Bob Johnson", Email: "bob@example.com"},
	}
	
	// Simple search filter
	if search != "" {
		filtered := []User{}
		for _, user := range users {
			if containsIgnoreCase(user.Name, search) || containsIgnoreCase(user.Email, search) {
				filtered = append(filtered, user)
			}
		}
		users = filtered
	}
	
	// Simple pagination
	start := (page - 1) * limit
	end := start + limit
	if start > len(users) {
		users = []User{}
	} else if end > len(users) {
		users = users[start:]
	} else {
		users = users[start:end]
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		sendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	// Simulate user lookup
	if userID == 1 {
		user := User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	
	sendError(w, "User not found", http.StatusNotFound)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		sendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// Simulate user update
	if userID == 1 {
		user := User{
			ID:    1,
			Name:  req.Name,
			Email: req.Email,
		}
		
		// Set defaults if not provided
		if user.Name == "" {
			user.Name = "John Doe"
		}
		if user.Email == "" {
			user.Email = "john@example.com"
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	
	sendError(w, "User not found", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		sendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	// Simulate user deletion
	if userID == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	
	sendError(w, "User not found", http.StatusNotFound)
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
		Code:  statusCode,
	})
}

func containsIgnoreCase(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}