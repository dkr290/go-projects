package auto

import (
	"net/http"
)

// New creates a new AutoSwagger instance (recommended).
func New(title, version string) *AutoSwagger {
	return NewAuto(title, version)
}

// NewSwaggerOnly creates a new basic Swagger instance without middleware.
func NewSwaggerOnly(title, version string) *Swagger {
	return NewSwagger(title, version)
}

// NewWithNetHTTP creates a new AutoSwagger instance with net/http adapter.
func NewWithNetHTTP(title, version string) (*AutoSwagger, *NetHTTPAdapter) {
	swagger := NewAuto(title, version)
	adapter := NewNetHTTPAdapter(swagger)
	return swagger, adapter
}

// NewWithNetHTTPMux creates a new AutoSwagger instance with net/http adapter and custom mux.
func NewWithNetHTTPMux(title, version string, mux *http.ServeMux) (*AutoSwagger, *NetHTTPAdapter) {
	swagger := NewAuto(title, version)
	adapter := NewNetHTTPAdapterWithMux(swagger, mux)
	return swagger, adapter
}

// NewWithChi creates a new AutoSwagger instance with chi adapter.
func NewWithChi(title, version string, router ChiRouter) (*AutoSwagger, *ChiAdapter) {
	swagger := NewAuto(title, version)
	adapter := NewChiAdapter(swagger, router)
	return swagger, adapter
}
