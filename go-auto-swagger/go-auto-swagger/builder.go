package auto

import (
	"net/http"
	"reflect"
)

// RouteBuilder helps build routes with documentation.
type RouteBuilder struct {
	swagger *AutoSwagger
	method  string
	path    string
	info    RouteInfo
}

// Route creates a new route builder.
func (as *AutoSwagger) Route(method, path string) *RouteBuilder {
	return &RouteBuilder{
		swagger: as,
		method:  method,
		path:    path,
		info: RouteInfo{
			Summary:     generateSummary(method, path),
			Description: generateDescription(method, path),
			Tags:        []string{extractTag(path)},
			OperationID: generateOperationID(method, path),
		},
	}
}

// GET creates a GET route builder.
func (as *AutoSwagger) GET(path string) *RouteBuilder {
	return as.Route("GET", path)
}

// POST creates a POST route builder.
func (as *AutoSwagger) POST(path string) *RouteBuilder {
	return as.Route("POST", path)
}

// PUT creates a PUT route builder.
func (as *AutoSwagger) PUT(path string) *RouteBuilder {
	return as.Route("PUT", path)
}

// DELETE creates a DELETE route builder.
func (as *AutoSwagger) DELETE(path string) *RouteBuilder {
	return as.Route("DELETE", path)
}

// PATCH creates a PATCH route builder.
func (as *AutoSwagger) PATCH(path string) *RouteBuilder {
	return as.Route("PATCH", path)
}

// Summary sets the summary for the route.
func (rb *RouteBuilder) Summary(summary string) *RouteBuilder {
	rb.info.Summary = summary
	return rb
}

// Description sets the description for the route.
func (rb *RouteBuilder) Description(description string) *RouteBuilder {
	rb.info.Description = description
	return rb
}

// Tags sets the tags for the route.
func (rb *RouteBuilder) Tags(tags ...string) *RouteBuilder {
	rb.info.Tags = tags
	return rb
}

// OperationID sets the operation ID for the route.
func (rb *RouteBuilder) OperationID(operationID string) *RouteBuilder {
	rb.info.OperationID = operationID
	return rb
}

// Request sets the request type for the route.
func (rb *RouteBuilder) Request(requestType interface{}) *RouteBuilder {
	rb.info.RequestType = requestType
	return rb
}

// Response sets the response type for the route.
func (rb *RouteBuilder) Response(responseType interface{}) *RouteBuilder {
	rb.info.ResponseType = responseType
	return rb
}

// Responses sets custom responses for the route.
func (rb *RouteBuilder) Responses(responses map[string]Response) *RouteBuilder {
	rb.info.Responses = responses
	return rb
}

// Parameter adds a parameter to the route.
func (rb *RouteBuilder) Parameter(param Parameter) *RouteBuilder {
	rb.info.Parameters = append(rb.info.Parameters, param)
	return rb
}

// Query adds a query parameter to the route.
func (rb *RouteBuilder) Query(name, description string, required bool, schema *Schema) *RouteBuilder {
	param := Parameter{
		Name:        name,
		In:          "query",
		Description: description,
		Required:    required,
		Schema:      schema,
	}
	return rb.Parameter(param)
}

// Path adds a path parameter to the route.
func (rb *RouteBuilder) Path(name, description string, schema *Schema) *RouteBuilder {
	param := Parameter{
		Name:        name,
		In:          "path",
		Description: description,
		Required:    true,
		Schema:      schema,
	}
	return rb.Parameter(param)
}

// Header adds a header parameter to the route.
func (rb *RouteBuilder) Header(name, description string, required bool, schema *Schema) *RouteBuilder {
	param := Parameter{
		Name:        name,
		In:          "header",
		Description: description,
		Required:    required,
		Schema:      schema,
	}
	return rb.Parameter(param)
}

// Register registers the route with the swagger documentation.
func (rb *RouteBuilder) Register() {
	rb.swagger.AddRoute(rb.method, rb.path, rb.info)
}

// RegisterHandler registers the route with the swagger documentation and returns a handler.
func (rb *RouteBuilder) RegisterHandler(handler http.HandlerFunc) *RouteBuilder {
	rb.Register()
	return rb
}

// RegisterHandlerWithTypes registers the route with inferred types from the handler.
func (rb *RouteBuilder) RegisterHandlerWithTypes(handler interface{}) {
	// Try to infer types from handler
	if rb.info.RequestType == nil || rb.info.ResponseType == nil {
		reqType, respType := rb.inferTypesFromHandler(handler)
		if rb.info.RequestType == nil {
			rb.info.RequestType = reqType
		}
		if rb.info.ResponseType == nil {
			rb.info.ResponseType = respType
		}
	}
	
	rb.Register()
}

// inferTypesFromHandler attempts to infer request and response types from handler.
func (rb *RouteBuilder) inferTypesFromHandler(handler interface{}) (interface{}, interface{}) {
	handlerType := reflect.TypeOf(handler)
	if handlerType == nil {
		return nil, nil
	}
	
	// For function types, analyze parameters and return types
	if handlerType.Kind() == reflect.Func {
		return rb.inferTypesFromFunc(handlerType)
	}
	
	return nil, nil
}

// inferTypesFromFunc infers types from function signature.
func (rb *RouteBuilder) inferTypesFromFunc(funcType reflect.Type) (interface{}, interface{}) {
	var requestType, responseType interface{}
	
	// Analyze function parameters
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		
		// Skip http.ResponseWriter and *http.Request
		if paramType == reflect.TypeOf((*http.ResponseWriter)(nil)).Elem() {
			continue
		}
		if paramType == reflect.TypeOf((*http.Request)(nil)) {
			continue
		}
		
		// Assume the first non-HTTP type is the request type
		if requestType == nil {
			requestType = reflect.New(paramType).Interface()
		}
	}
	
	// Analyze function return types
	for i := 0; i < funcType.NumOut(); i++ {
		returnType := funcType.Out(i)
		
		// Skip error type
		if returnType == reflect.TypeOf((*error)(nil)).Elem() {
			continue
		}
		
		// Assume the first non-error type is the response type
		if responseType == nil {
			responseType = reflect.New(returnType).Interface()
		}
	}
	
	return requestType, responseType
}

// Common schema helpers
var (
	StringSchema = &Schema{Type: "string"}
	IntSchema    = &Schema{Type: "integer"}
	BoolSchema   = &Schema{Type: "boolean"}
	NumberSchema = &Schema{Type: "number"}
)

// PathParam creates a path parameter with string type.
func PathParam(name, description string) Parameter {
	return Parameter{
		Name:        name,
		In:          "path",
		Description: description,
		Required:    true,
		Schema:      StringSchema,
	}
}

// QueryParam creates a query parameter with string type.
func QueryParam(name, description string, required bool) Parameter {
	return Parameter{
		Name:        name,
		In:          "query",
		Description: description,
		Required:    required,
		Schema:      StringSchema,
	}
}

// HeaderParam creates a header parameter with string type.
func HeaderParam(name, description string, required bool) Parameter {
	return Parameter{
		Name:        name,
		In:          "header",
		Description: description,
		Required:    required,
		Schema:      StringSchema,
	}
}

// JSONResponse creates a JSON response.
func JSONResponse(description string, responseType interface{}) Response {
	schema := &Schema{Type: "object"}
	if responseType != nil {
		t := reflect.TypeOf(responseType)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		schema.Ref = "#/components/schemas/" + t.Name()
	}
	
	return Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: schema,
			},
		},
	}
}

// ErrorResponse creates a standard error response.
func ErrorResponse(description string) Response {
	return Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{
					Type: "object",
					Properties: map[string]*Schema{
						"error": {Type: "string"},
						"code":  {Type: "integer"},
					},
				},
			},
		},
	}
}

// DefaultResponses creates default response map.
func DefaultResponses(successType interface{}) map[string]Response {
	responses := map[string]Response{
		"400": ErrorResponse("Bad Request"),
		"401": ErrorResponse("Unauthorized"),
		"403": ErrorResponse("Forbidden"),
		"404": ErrorResponse("Not Found"),
		"500": ErrorResponse("Internal Server Error"),
	}
	
	if successType != nil {
		responses["200"] = JSONResponse("Success", successType)
	}
	
	return responses
}