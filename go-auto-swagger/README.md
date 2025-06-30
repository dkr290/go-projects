# Go Auto Swagger

🚀 **Automatically generate OpenAPI 3.0 (Swagger) documentation for your Go APIs**

Go Auto Swagger is a lightweight, powerful library that automatically generates OpenAPI documentation from your Go API code. No more manual YAML writing or maintaining separate documentation files!

## ✨ Features

- 🔄 **Automatic Schema Generation** - Converts Go structs to OpenAPI schemas
- 🌐 **Multiple Router Support** - Works with `net/http`, Chi router, and extensible to others
- ⛓️ **Fluent API** - Chainable methods for easy route documentation
- 🧠 **Type Inference** - Automatically infers request/response types from handlers
- 📖 **Built-in Swagger UI** - Serves interactive documentation out of the box
- 🔧 **Parameter Support** - Query, path, header parameters with validation
- 📝 **Custom Responses** - Define multiple response codes and descriptions
- 🏷️ **Tags & Grouping** - Organize endpoints logically
- 📊 **JSON Tags Support** - Respects struct tags for field naming and validation

## 📦 Installation

```bash
go get github.com/dani/go-auto-swagger/go-auto-swagger
```

## 🚀 Quick Start

### Basic Example with net/http

```go
package main

import (
    "encoding/json"
    "net/http"
    "log"

    auto "github.com/dani/go-auto-swagger/go-auto-swagger"
)

type User struct {
    ID    int    `json:"id" description:"User ID" example:"1"`
    Name  string `json:"name" description:"User name" example:"John Doe"`
    Email string `json:"email" description:"User email" example:"john@example.com"`
}

type CreateUserRequest struct {
    Name  string `json:"name" description:"User name"`
    Email string `json:"email" description:"User email"`
}

func main() {
    // Create swagger instance with net/http adapter
    swagger, adapter := auto.NewWithNetHTTP("User API", "1.0.0")

    // Configure API info
    swagger.SetInfo(auto.Info{
        Title:       "User Management API",
        Description: "A simple API for managing users",
        Version:     "1.0.0",
        Contact: auto.Contact{
            Name:  "API Support",
            Email: "support@example.com",
        },
    })

    // Add server info
    swagger.AddServer("http://localhost:8080", "Development server")
    swagger.AddTag("users", "User management operations")

    // Register endpoints with automatic documentation
    adapter.GET("/users", listUsers).
        Summary("List Users").
        Description("Get all users with optional pagination").
        Tags("users").
        Query("limit", "Number of users to return", false, auto.IntSchema).
        Query("offset", "Number of users to skip", false, auto.IntSchema).
        Response(&[]User{})

    adapter.POST("/users", createUser).
        Summary("Create User").
        Description("Create a new user").
        Tags("users").
        Request(&CreateUserRequest{}).
        Response(&User{})

    adapter.GET("/users/{id}", getUser).
        Summary("Get User").
        Description("Get a user by ID").
        Tags("users").
        Path("id", "User ID", auto.IntSchema).
        Response(&User{})

    // Serve documentation
    adapter.ServeDocs("/swagger.json")
    adapter.ServeSwaggerUI("/docs", "/swagger.json")

    log.Println("Server starting on :8080")
    log.Println("API Documentation: http://localhost:8080/docs")
    log.Println("Swagger JSON: http://localhost:8080/swagger.json")

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "John Doe", Email: "john@example.com"},
        {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user := User{
        ID:    3,
        Name:  req.Name,
        Email: req.Email,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL path
    user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

Run the server and visit:

- **Interactive Documentation**: http://localhost:8080/docs
- **Raw Swagger JSON**: http://localhost:8080/swagger.json

## 🔌 Router Adapters

### Chi Router

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    auto "github.com/dani/go-auto-swagger/go-auto-swagger"
)

func main() {
    r := chi.NewRouter()
    swagger := auto.New("My Chi API", "1.0.0")
    adapter := auto.NewChiAdapter(swagger, r)

    // Register routes with documentation
    adapter.GET("/health", healthCheck).
        Summary("Health Check").
        Description("Check if the API is running").
        Tags("system").
        Response(map[string]string{"status": "ok"})

    adapter.GET("/users/{id}", getUser).
        Summary("Get User").
        Description("Retrieve a user by ID").
        Tags("users").
        Path("id", "User ID", auto.IntSchema).
        Response(&User{})

    // Serve documentation
    adapter.ServeDocs("/swagger.json")
    adapter.ServeSwaggerUI("/docs", "/swagger.json")

    http.ListenAndServe(":8080", r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"status": "ok"}`))
}
```

## 📚 Advanced Usage

### Custom Response Codes

```go
adapter.POST("/users", createUser).
    Summary("Create User").
    Request(&CreateUserRequest{}).
    Responses(map[string]auto.Response{
        "201": auto.JSONResponse("User created successfully", &User{}),
        "400": auto.ErrorResponse("Invalid request data"),
        "409": auto.ErrorResponse("User already exists"),
        "500": auto.ErrorResponse("Internal server error"),
    })
```

### Multiple Parameters

```go
adapter.GET("/users", listUsers).
    Summary("List Users").
    Query("search", "Search term", false, auto.StringSchema).
    Query("page", "Page number", false, auto.IntSchema).
    Query("limit", "Items per page", false, auto.IntSchema).
    Header("X-API-Key", "API Key", true, auto.StringSchema).
    Response(&[]User{})
```

### Complex Types

```go
type UserWithMetadata struct {
    User
    CreatedAt time.Time            `json:"created_at" description:"Creation timestamp"`
    UpdatedAt time.Time            `json:"updated_at" description:"Last update timestamp"`
    Tags      []string             `json:"tags" description:"User tags"`
    Metadata  map[string]interface{} `json:"metadata" description:"Additional metadata"`
}

adapter.GET("/users/{id}/full", getUserWithMetadata).
    Summary("Get User with Metadata").
    Path("id", "User ID", auto.IntSchema).
    Response(&UserWithMetadata{})
```

### Route Builder Pattern

```go
// You can also build routes step by step
userRoutes := swagger.Route("GET", "/users/{id}")
userRoutes.Summary("Get User by ID")
userRoutes.Description("Retrieves a user by their unique identifier")
userRoutes.Tags("users")
userRoutes.Path("id", "User ID", auto.IntSchema)
userRoutes.Response(&User{})
userRoutes.Register()
```

## 🎨 Schema Types

Go Auto Swagger provides built-in schema helpers:

```go
auto.StringSchema  // String type
auto.IntSchema     // Integer type
auto.BoolSchema    // Boolean type
auto.NumberSchema  // Number (float) type

// Custom schema
customSchema := &auto.Schema{
    Type:        "string",
    Format:      "email",
    Description: "Email address",
    Example:     "user@example.com",
}
```

## 🏷️ Struct Tags

Enhance your documentation with struct tags:

```go
type User struct {
    ID       int       `json:"id" description:"Unique identifier" example:"123"`
    Name     string    `json:"name" description:"Full name" example:"John Doe"`
    Email    string    `json:"email" description:"Email address" example:"john@example.com"`
    Age      int       `json:"age,omitempty" description:"User age" example:"30"`
    IsActive bool      `json:"is_active" description:"Account status" example:"true"`
    Tags     []string  `json:"tags,omitempty" description:"User tags"`
}
```

Supported tags:

- `json:"field_name"` - Field name in JSON
- `json:",omitempty"` - Optional field
- `description:"..."` - Field description
- `example:"..."` - Example value

## 🌟 API Reference

### Creating Instances

```go
// Basic swagger instance
swagger := auto.New("API Title", "1.0.0")

// With net/http adapter
swagger, adapter := auto.NewWithNetHTTP("API Title", "1.0.0")

// With chi adapter
adapter := auto.NewChiAdapter(swagger, chiRouter)
```

### Configuration

```go
// Set API information
swagger.SetInfo(auto.Info{
    Title:          "My API",
    Description:    "API description",
    Version:        "1.0.0",
    TermsOfService: "https://example.com/terms",
    Contact: auto.Contact{
        Name:  "Support",
        Email: "support@example.com",
        URL:   "https://example.com/support",
    },
    License: auto.License{
        Name: "MIT",
        URL:  "https://opensource.org/licenses/MIT",
    },
})

// Add servers
swagger.AddServer("https://api.example.com", "Production")
swagger.AddServer("https://staging-api.example.com", "Staging")

// Add tags
swagger.AddTag("users", "User management operations")
swagger.AddTag("auth", "Authentication operations")
```

### Route Builder Methods

```go
builder := adapter.GET("/path", handler)

// Documentation
builder.Summary("Short summary")
builder.Description("Detailed description")
builder.Tags("tag1", "tag2")
builder.OperationID("uniqueOperationId")

// Parameters
builder.Path("id", "Description", schema)
builder.Query("param", "Description", required, schema)
builder.Header("X-Header", "Description", required, schema)

// Request/Response
builder.Request(&RequestType{})
builder.Response(&ResponseType{})
builder.Responses(map[string]auto.Response{...})
```

## 🔧 Extending

### Custom Adapters

You can create adapters for other routers:

```go
type MyRouterAdapter struct {
    swagger *auto.AutoSwagger
    router  MyCustomRouter
}

func (a *MyRouterAdapter) GET(path string, handler http.HandlerFunc) *auto.RouteBuilder {
    // Register with your router
    a.router.Get(path, handler)

    // Return swagger route builder
    return a.swagger.GET(path)
}
```

## 📝 Examples

Check out the complete examples in the repository:

- [Basic net/http example](examples/main.go)
- [Chi router example](examples/chi/main.go)
- [Advanced net/http example](examples/nethttp/main.go)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 🙏 Acknowledgments

- OpenAPI 3.0 Specification
- Swagger UI for the interactive documentation interface
- Go community for inspiration and best practices

## 📞 Support

If you have any questions or need help:

- 📧 Create an issue on GitHub
- 💬 Start a discussion in the repository
- 📖 Check the examples directory

---

**Made with ❤️ for the Go community**

