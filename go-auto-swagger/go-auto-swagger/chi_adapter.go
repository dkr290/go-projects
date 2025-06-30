package auto

import (
	"net/http"
)

// ChiAdapter provides integration with chi router.
// Note: This assumes chi router interface. Install with: go get github.com/go-chi/chi/v5
type ChiAdapter struct {
	swagger *AutoSwagger
	router  ChiRouter
}

// ChiRouter interface to avoid direct dependency on chi
type ChiRouter interface {
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Put(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
	Patch(pattern string, handlerFn http.HandlerFunc)
	Options(pattern string, handlerFn http.HandlerFunc)
	Head(pattern string, handlerFn http.HandlerFunc)
	HandleFunc(pattern string, handlerFn http.HandlerFunc)
	Handle(pattern string, handler http.Handler)
	Mount(pattern string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// NewChiAdapter creates a new chi adapter.
func NewChiAdapter(swagger *AutoSwagger, router ChiRouter) *ChiAdapter {
	return &ChiAdapter{
		swagger: swagger,
		router:  router,
	}
}

// GET registers a GET handler with automatic documentation.
func (adapter *ChiAdapter) GET(pattern string, handler http.HandlerFunc) *RouteBuilder {
	// Register with chi router
	adapter.router.Get(pattern, handler)
	
	// Return route builder for documentation
	return adapter.swagger.GET(pattern)
}

// POST registers a POST handler with automatic documentation.
func (adapter *ChiAdapter) POST(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Post(pattern, handler)
	return adapter.swagger.POST(pattern)
}

// PUT registers a PUT handler with automatic documentation.
func (adapter *ChiAdapter) PUT(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Put(pattern, handler)
	return adapter.swagger.PUT(pattern)
}

// DELETE registers a DELETE handler with automatic documentation.
func (adapter *ChiAdapter) DELETE(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Delete(pattern, handler)
	return adapter.swagger.DELETE(pattern)
}

// PATCH registers a PATCH handler with automatic documentation.
func (adapter *ChiAdapter) PATCH(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Patch(pattern, handler)
	return adapter.swagger.PATCH(pattern)
}

// OPTIONS registers an OPTIONS handler with automatic documentation.
func (adapter *ChiAdapter) OPTIONS(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Options(pattern, handler)
	return adapter.swagger.Route("OPTIONS", pattern)
}

// HEAD registers a HEAD handler with automatic documentation.
func (adapter *ChiAdapter) HEAD(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.router.Head(pattern, handler)
	return adapter.swagger.Route("HEAD", pattern)
}

// HandleFunc registers a handler function with automatic documentation.
func (adapter *ChiAdapter) HandleFunc(pattern string, handler http.HandlerFunc) {
	adapter.router.HandleFunc(pattern, handler)
	
	// Add basic documentation for generic handler
	info := adapter.swagger.InferFromHandler("GET", pattern, handler)
	adapter.swagger.AddRoute("GET", pattern, info)
}

// Handle registers a handler with automatic documentation.
func (adapter *ChiAdapter) Handle(pattern string, handler http.Handler) {
	adapter.router.Handle(pattern, handler)
	
	// Add basic documentation for generic handler
	info := RouteInfo{
		Summary:     "Handle " + pattern,
		Description: "Generic handler for " + pattern,
		Tags:        []string{extractTag(pattern)},
	}
	adapter.swagger.AddRoute("GET", pattern, info)
}

// Mount mounts a sub-router with automatic documentation.
func (adapter *ChiAdapter) Mount(pattern string, handler http.Handler) {
	adapter.router.Mount(pattern, handler)
	
	// Add basic documentation for mounted handler
	info := RouteInfo{
		Summary:     "Mount " + pattern,
		Description: "Mounted sub-router at " + pattern,
		Tags:        []string{extractTag(pattern)},
	}
	adapter.swagger.AddRoute("GET", pattern+"/*", info)
}

// ServeHTTP implements http.Handler interface.
func (adapter *ChiAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	adapter.router.ServeHTTP(w, r)
}

// ServeDocs serves the Swagger documentation at the specified pattern.
func (adapter *ChiAdapter) ServeDocs(pattern string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		json, err := adapter.swagger.ToJSON()
		if err != nil {
			http.Error(w, "Failed to generate documentation", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
	
	adapter.router.Get(pattern, handler)
}

// ServeSwaggerUI serves the Swagger UI at the specified pattern.
func (adapter *ChiAdapter) ServeSwaggerUI(pattern, docsURL string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		html := generateSwaggerUIHTML(docsURL)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
	
	adapter.router.Get(pattern, handler)
}

// Group creates a route group with common documentation properties.
func (adapter *ChiAdapter) Group(fn func(ChiGroupAdapter)) {
	// This would need to be implemented with chi's Group functionality
	// For now, we provide a basic grouping interface
	group := &ChiGroupAdapter{
		adapter: adapter,
		prefix:  "",
		tags:    []string{},
	}
	fn(*group)
}

// ChiGroupAdapter provides grouped route registration with shared documentation.
type ChiGroupAdapter struct {
	adapter *ChiAdapter
	prefix  string
	tags    []string
}

// WithTags sets default tags for the group.
func (group *ChiGroupAdapter) WithTags(tags ...string) *ChiGroupAdapter {
	group.tags = tags
	return group
}

// WithPrefix sets a prefix for the group.
func (group *ChiGroupAdapter) WithPrefix(prefix string) *ChiGroupAdapter {
	group.prefix = prefix
	return group
}

// GET registers a GET handler in the group.
func (group *ChiGroupAdapter) GET(pattern string, handler http.HandlerFunc) *RouteBuilder {
	fullPattern := group.prefix + pattern
	return group.adapter.GET(fullPattern, handler).Tags(group.tags...)
}

// POST registers a POST handler in the group.
func (group *ChiGroupAdapter) POST(pattern string, handler http.HandlerFunc) *RouteBuilder {
	fullPattern := group.prefix + pattern
	return group.adapter.POST(fullPattern, handler).Tags(group.tags...)
}

// PUT registers a PUT handler in the group.
func (group *ChiGroupAdapter) PUT(pattern string, handler http.HandlerFunc) *RouteBuilder {
	fullPattern := group.prefix + pattern
	return group.adapter.PUT(fullPattern, handler).Tags(group.tags...)
}

// DELETE registers a DELETE handler in the group.
func (group *ChiGroupAdapter) DELETE(pattern string, handler http.HandlerFunc) *RouteBuilder {
	fullPattern := group.prefix + pattern
	return group.adapter.DELETE(fullPattern, handler).Tags(group.tags...)
}

// PATCH registers a PATCH handler in the group.
func (group *ChiGroupAdapter) PATCH(pattern string, handler http.HandlerFunc) *RouteBuilder {
	fullPattern := group.prefix + pattern
	return group.adapter.PATCH(fullPattern, handler).Tags(group.tags...)
}