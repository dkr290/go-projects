package auto

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

// AutoSwagger wraps a Swagger instance with automatic route discovery.
type AutoSwagger struct {
	*Swagger
	middlewareEnabled bool
}

// NewAuto creates a new AutoSwagger instance.
func NewAuto(title, version string) *AutoSwagger {
	return &AutoSwagger{
		Swagger:           NewSwagger(title, version),
		middlewareEnabled: true,
	}
}

// EnableMiddleware enables automatic route discovery.
func (as *AutoSwagger) EnableMiddleware() {
	as.middlewareEnabled = true
}

// DisableMiddleware disables automatic route discovery.
func (as *AutoSwagger) DisableMiddleware() {
	as.middlewareEnabled = false
}

// Handler creates a middleware handler that automatically documents routes.
func (as *AutoSwagger) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if as.middlewareEnabled {
			as.documentRequest(r)
		}
		next.ServeHTTP(w, r)
	})
}

// HandlerFunc creates a middleware handler function.
func (as *AutoSwagger) HandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if as.middlewareEnabled {
			as.documentRequest(r)
		}
		next(w, r)
	}
}

// documentRequest automatically documents the request.
func (as *AutoSwagger) documentRequest(r *http.Request) {
	path := r.URL.Path
	method := r.Method
	
	// Check if route already exists
	if pathItem, exists := as.Paths[path]; exists {
		if as.operationExists(pathItem, method) {
			return
		}
	}
	
	// Create basic route info
	info := RouteInfo{
		Summary:     generateSummary(method, path),
		Description: generateDescription(method, path),
		Tags:        []string{extractTag(path)},
	}
	
	as.AddRoute(method, path, info)
}

// operationExists checks if an operation already exists for the given method.
func (as *AutoSwagger) operationExists(pathItem PathItem, method string) bool {
	switch strings.ToUpper(method) {
	case "GET":
		return pathItem.Get != nil
	case "POST":
		return pathItem.Post != nil
	case "PUT":
		return pathItem.Put != nil
	case "DELETE":
		return pathItem.Delete != nil
	case "PATCH":
		return pathItem.Patch != nil
	}
	return false
}

// generateSummary generates a summary for the route.
func generateSummary(method, path string) string {
	action := methodToAction(method)
	resource := extractResource(path)
	return action + " " + resource
}

// generateDescription generates a description for the route.
func generateDescription(method, path string) string {
	action := methodToAction(method)
	resource := extractResource(path)
	return action + " " + resource + " endpoint"
}

// methodToAction converts HTTP method to action verb.
func methodToAction(method string) string {
	switch strings.ToUpper(method) {
	case "GET":
		return "Get"
	case "POST":
		return "Create"
	case "PUT":
		return "Update"
	case "DELETE":
		return "Delete"
	case "PATCH":
		return "Patch"
	default:
		return strings.Title(strings.ToLower(method))
	}
}

// extractResource extracts resource name from path.
func extractResource(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return "Resource"
	}
	
	// Get the last part that doesn't look like an ID
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		// Skip parts that look like IDs (numbers, UUIDs, etc.)
		if !isLikelyID(part) {
			return strings.Title(part)
		}
	}
	
	return strings.Title(parts[0])
}

// extractTag extracts tag from path.
func extractTag(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 {
		return "default"
	}
	return parts[0]
}

// isLikelyID checks if a path segment looks like an ID.
func isLikelyID(segment string) bool {
	// Check for common ID patterns
	if segment == "" {
		return false
	}
	
	// Check for parameter patterns like {id}, :id
	if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
		return true
	}
	if strings.HasPrefix(segment, ":") {
		return true
	}
	
	// Check if it's all digits
	for _, r := range segment {
		if r < '0' || r > '9' {
			return false
		}
	}
	
	return len(segment) > 0
}

// Route represents a documented route with its handler.
type Route struct {
	Method      string
	Path        string
	Handler     interface{}
	Info        RouteInfo
}

// RegisterRoute registers a route with detailed information.
func (as *AutoSwagger) RegisterRoute(method, path string, handler interface{}, info RouteInfo) {
	// Add the route to swagger documentation
	as.AddRoute(method, path, info)
	
	// Store route information for later use
	as.storeRouteInfo(method, path, handler, info)
}

// storeRouteInfo stores route information for introspection.
func (as *AutoSwagger) storeRouteInfo(method, path string, handler interface{}, info RouteInfo) {
	// This could be extended to store route metadata for further processing
	// For now, we just ensure the route is documented
}

// InferFromHandler attempts to infer route information from handler function.
func (as *AutoSwagger) InferFromHandler(method, path string, handler interface{}) RouteInfo {
	info := RouteInfo{
		Summary:     generateSummary(method, path),
		Description: generateDescription(method, path),
		Tags:        []string{extractTag(path)},
		OperationID: generateOperationID(method, path),
	}
	
	// Try to infer types from handler signature
	if handlerType := reflect.TypeOf(handler); handlerType != nil {
		info.RequestType, info.ResponseType = as.inferTypesFromHandler(handlerType)
	}
	
	return info
}

// inferTypesFromHandler attempts to infer request/response types from handler.
func (as *AutoSwagger) inferTypesFromHandler(handlerType reflect.Type) (interface{}, interface{}) {
	// This is a simplified implementation
	// In a real implementation, you'd analyze the handler signature more thoroughly
	return nil, nil
}

// generateOperationID generates an operation ID for the route.
func generateOperationID(method, path string) string {
	// Clean the path and method to create a valid operation ID
	cleanPath := strings.ReplaceAll(strings.Trim(path, "/"), "/", "_")
	cleanPath = strings.ReplaceAll(cleanPath, "{", "")
	cleanPath = strings.ReplaceAll(cleanPath, "}", "")
	cleanPath = strings.ReplaceAll(cleanPath, ":", "")
	
	return strings.ToLower(method) + "_" + cleanPath
}

// GetFunctionName gets the name of a function.
func GetFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}