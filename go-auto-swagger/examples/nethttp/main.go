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
	// Create a new swagger instance
	swagger := auto.New("User Management API", "1.0.0")
	
	// Configure API metadata
	swagger.SetInfo(auto.Info{
		Title:       "User Management API",
		Description: "A RESTful API for managing users built with net/http",
		Version:     "1.0.0",
		Contact: auto.Contact{
			Name:  "API Team",
			Email: "api@example.com",
		},
	})
	
	// Add server information
	swagger.AddServer("http://localhost:8080", "Development server")
	
	// Add tags for organization
	swagger.AddTag("users", "User management endpoints")
	swagger.AddTag("health", "Health check endpoints")

	// Register POST /users endpoint
	postUsersInfo := auto.RouteInfo{
		Summary:     "Create a new user",
		Description: "Creates a new user with the provided name and email",
		Tags:        []string{"users"},
		OperationID: "createUser",
		RequestType: &CreateUserRequest{},
		ResponseType: &CreateUserResponse{},
		Responses: map[string]auto.Response{
			"201": auto.JSONResponse("User created successfully", &CreateUserResponse{}),
			"400": auto.ErrorResponse("Invalid request body"),
			"500": auto.ErrorResponse("Internal server error"),
		},
	}
	swagger.AddRoute("POST", "/users", postUsersInfo)

	// Register GET /users/{id} endpoint
	getUserInfo := auto.RouteInfo{
		Summary:     "Get user by ID",
		Description: "Retrieves a user by their unique identifier",
		Tags:        []string{"users"},
		OperationID: "getUserById",
		Parameters: []auto.Parameter{
			auto.PathParam("id", "User ID"),
		},
		ResponseType: &User{},
		Responses: map[string]auto.Response{
			"200": auto.JSONResponse("User found", &User{}),
			"404": auto.ErrorResponse("User not found"),
			"500": auto.ErrorResponse("Internal server error"),
		},
	}
	swagger.AddRoute("GET", "/users/{id}", getUserInfo)
	
	// Register GET /health endpoint
	healthInfo := auto.RouteInfo{
		Summary:     "Health check",
		Description: "Returns the health status of the API",
		Tags:        []string{"health"},
		OperationID: "healthCheck",
		ResponseType: map[string]interface{}{"status": "ok"},
		Responses: map[string]auto.Response{
			"200": auto.JSONResponse("API is healthy", map[string]interface{}{"status": "ok"}),
		},
	}
	swagger.AddRoute("GET", "/health", healthInfo)
	
	// Set up HTTP handlers
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		}
	})
	
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		}
	})
	
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Serve Swagger JSON documentation
	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swaggerJSON, err := swagger.ToJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate documentation"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerJSON)
	})
	
	// Serve Swagger UI
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
<html>
<head>
    <title>API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui.css" />
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '/swagger.json',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
        };
    </script>
</body>
</html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	log.Println("Server starting on :8080")
	log.Println("API Documentation: http://localhost:8080/docs")
	log.Println("Swagger JSON: http://localhost:8080/swagger.json")
	log.Println("Health Check: http://localhost:8080/health")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Simulate user creation
	user := User{
		ID:   1,
		Name: req.Name,
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

func getUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL (simplified)
	// In a real application, you'd parse the ID from the URL path
	
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
