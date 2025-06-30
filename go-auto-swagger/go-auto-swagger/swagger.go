package auto

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Swagger represents the OpenAPI specification.
type Swagger struct {
	OpenAPI    string                    `json:"openapi"`
	Info       Info                      `json:"info"`
	Servers    []Server                  `json:"servers,omitempty"`
	Paths      map[string]PathItem       `json:"paths"`
	Components Components                `json:"components"`
	Tags       []Tag                     `json:"tags,omitempty"`
}

// Info represents the metadata of the API.
type Info struct {
	Title          string  `json:"title"`
	Description    string  `json:"description,omitempty"`
	Version        string  `json:"version"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
}

// Contact represents contact information.
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License represents license information.
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

// Server represents server information.
type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

// Tag represents a tag for grouping operations.
type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// PathItem represents a path and its operations.
type PathItem struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
}

// Operation represents an HTTP operation.
type Operation struct {
	Tags        []string              `json:"tags,omitempty"`
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty"`
	Responses   map[string]Response   `json:"responses"`
	Security    []map[string][]string `json:"security,omitempty"`
}

// Parameter represents a parameter.
type Parameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

// RequestBody represents a request body.
type RequestBody struct {
	Description string                       `json:"description,omitempty"`
	Content     map[string]MediaType         `json:"content"`
	Required    bool                         `json:"required,omitempty"`
}

// Response represents a response.
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Headers     map[string]Header    `json:"headers,omitempty"`
}

// MediaType represents a media type.
type MediaType struct {
	Schema *Schema `json:"schema,omitempty"`
}

// Header represents a header.
type Header struct {
	Description string  `json:"description,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

// Components represents the reusable components of the API.
type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
}

// SecurityScheme represents a security scheme.
type SecurityScheme struct {
	Type         string `json:"type"`
	Description  string `json:"description,omitempty"`
	Name         string `json:"name,omitempty"`
	In           string `json:"in,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
}

// Schema represents the schema of a data type.
type Schema struct {
	Type                 string             `json:"type,omitempty"`
	Format               string             `json:"format,omitempty"`
	Description          string             `json:"description,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	Required             []string           `json:"required,omitempty"`
	Enum                 []interface{}      `json:"enum,omitempty"`
	Example              interface{}        `json:"example,omitempty"`
	Ref                  string             `json:"$ref,omitempty"`
	AdditionalProperties interface{}        `json:"additionalProperties,omitempty"`
	Minimum              *float64           `json:"minimum,omitempty"`
	Maximum              *float64           `json:"maximum,omitempty"`
	MinLength            *int               `json:"minLength,omitempty"`
	MaxLength            *int               `json:"maxLength,omitempty"`
}

// NewSwagger creates a new Swagger instance.
func NewSwagger(title, version string) *Swagger {
	return &Swagger{
		OpenAPI: "3.0.0",
		Info: Info{
			Title:   title,
			Version: version,
		},
		Paths: make(map[string]PathItem),
		Components: Components{
			Schemas:         make(map[string]Schema),
			Responses:       make(map[string]Response),
			Parameters:      make(map[string]Parameter),
			RequestBodies:   make(map[string]RequestBody),
			Headers:         make(map[string]Header),
			SecuritySchemes: make(map[string]SecurityScheme),
		},
	}
}

// SetInfo sets the API info.
func (s *Swagger) SetInfo(info Info) {
	s.Info = info
}

// AddServer adds a server to the API documentation.
func (s *Swagger) AddServer(url, description string) {
	s.Servers = append(s.Servers, Server{
		URL:         url,
		Description: description,
	})
}

// AddTag adds a tag to the API documentation.
func (s *Swagger) AddTag(name, description string) {
	s.Tags = append(s.Tags, Tag{
		Name:        name,
		Description: description,
	})
}

// RouteInfo contains metadata for a route.
type RouteInfo struct {
	Summary      string
	Description  string
	Tags         []string
	OperationID  string
	RequestType  interface{}
	ResponseType interface{}
	Parameters   []Parameter
	Responses    map[string]Response
}

// AddRoute adds a new route to the Swagger documentation with metadata.
func (s *Swagger) AddRoute(method, path string, info RouteInfo) {
	// Add schemas for request and response types
	if info.RequestType != nil {
		s.addSchemaFromType(reflect.TypeOf(info.RequestType))
	}
	if info.ResponseType != nil {
		s.addSchemaFromType(reflect.TypeOf(info.ResponseType))
	}

	// Get or create path item
	pathItem, exists := s.Paths[path]
	if !exists {
		pathItem = PathItem{}
	}

	// Create operation
	operation := &Operation{
		Tags:        info.Tags,
		Summary:     info.Summary,
		Description: info.Description,
		OperationID: info.OperationID,
		Parameters:  info.Parameters,
		Responses:   info.Responses,
	}

	// Add request body if request type is provided
	if info.RequestType != nil {
		operation.RequestBody = s.createRequestBody(info.RequestType)
	}

	// Add default response if not provided
	if operation.Responses == nil {
		operation.Responses = make(map[string]Response)
	}
	if _, exists := operation.Responses["200"]; !exists && info.ResponseType != nil {
		operation.Responses["200"] = s.createResponse("Success", info.ResponseType)
	}

	// Set operation on path item
	switch strings.ToUpper(method) {
	case "GET":
		pathItem.Get = operation
	case "POST":
		pathItem.Post = operation
	case "PUT":
		pathItem.Put = operation
	case "DELETE":
		pathItem.Delete = operation
	case "PATCH":
		pathItem.Patch = operation
	}

	s.Paths[path] = pathItem
}

func (s *Swagger) createRequestBody(requestType interface{}) *RequestBody {
	t := reflect.TypeOf(requestType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	return &RequestBody{
		Description: "Request body",
		Required:    true,
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{
					Ref: "#/components/schemas/" + t.Name(),
				},
			},
		},
	}
}

func (s *Swagger) createResponse(description string, responseType interface{}) Response {
	t := reflect.TypeOf(responseType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	return Response{
		Description: description,
		Content: map[string]MediaType{
			"application/json": {
				Schema: &Schema{
					Ref: "#/components/schemas/" + t.Name(),
				},
			},
		},
	}
}

func (s *Swagger) addSchemaFromType(t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	if _, ok := s.Components.Schemas[t.Name()]; ok {
		return
	}

	schema := s.generateSchemaFromType(t)
	s.Components.Schemas[t.Name()] = schema
}

func (s *Swagger) generateSchemaFromType(t reflect.Type) Schema {
	switch t.Kind() {
	case reflect.String:
		return Schema{Type: "string"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Schema{Type: "integer"}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Schema{Type: "integer", Minimum: &[]float64{0}[0]}
	case reflect.Float32, reflect.Float64:
		return Schema{Type: "number"}
	case reflect.Bool:
		return Schema{Type: "boolean"}
	case reflect.Slice, reflect.Array:
		itemSchema := s.generateSchemaFromType(t.Elem())
		return Schema{
			Type:  "array",
			Items: &itemSchema,
		}
	case reflect.Map:
		return Schema{
			Type: "object",
			AdditionalProperties: true,
		}
	case reflect.Struct:
		return s.generateStructSchema(t)
	case reflect.Ptr:
		return s.generateSchemaFromType(t.Elem())
	default:
		return Schema{Type: "object"}
	}
}

func (s *Swagger) generateStructSchema(t reflect.Type) Schema {
	schema := Schema{
		Type:       "object",
		Properties: make(map[string]*Schema),
		Required:   []string{},
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
		// Skip unexported fields
		if !field.IsExported() {
			continue
		}
		
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = strings.ToLower(field.Name)
		}
		
		// Parse json tag
		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]
		
		// Skip if json:"-"
		if fieldName == "-" {
			continue
		}

		// Generate schema for field
		fieldSchema := s.generateSchemaFromType(field.Type)
		
		// Add description from struct tag
		if desc := field.Tag.Get("description"); desc != "" {
			fieldSchema.Description = desc
		}
		
		// Add example from struct tag
		if example := field.Tag.Get("example"); example != "" {
			fieldSchema.Example = example
		}
		
		// Check if field is required (no omitempty tag)
		isRequired := true
		for _, part := range tagParts[1:] {
			if part == "omitempty" {
				isRequired = false
				break
			}
		}
		
		if isRequired {
			schema.Required = append(schema.Required, fieldName)
		}

		schema.Properties[fieldName] = &fieldSchema
		
		// Handle nested structs
		if field.Type.Kind() == reflect.Struct {
			s.addSchemaFromType(field.Type)
		} else if field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.Struct {
			s.addSchemaFromType(field.Type.Elem())
		} else if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			s.addSchemaFromType(field.Type.Elem())
		}
	}

	return schema
}

// ToJSON returns the Swagger documentation in JSON format.
func (s *Swagger) ToJSON() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}