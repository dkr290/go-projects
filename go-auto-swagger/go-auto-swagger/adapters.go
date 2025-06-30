package auto

import (
	"net/http"
)

// NetHTTPAdapter provides integration with net/http.
type NetHTTPAdapter struct {
	swagger *AutoSwagger
	mux     *http.ServeMux
}

// NewNetHTTPAdapter creates a new net/http adapter.
func NewNetHTTPAdapter(swagger *AutoSwagger) *NetHTTPAdapter {
	return &NetHTTPAdapter{
		swagger: swagger,
		mux:     http.NewServeMux(),
	}
}

// NewNetHTTPAdapterWithMux creates a new net/http adapter with existing mux.
func NewNetHTTPAdapterWithMux(swagger *AutoSwagger, mux *http.ServeMux) *NetHTTPAdapter {
	return &NetHTTPAdapter{
		swagger: swagger,
		mux:     mux,
	}
}

// HandleFunc registers a handler function with automatic documentation.
func (adapter *NetHTTPAdapter) HandleFunc(pattern string, handler http.HandlerFunc) {
	// Register with automatic documentation
	info := adapter.swagger.InferFromHandler("GET", pattern, handler)
	adapter.swagger.AddRoute("GET", pattern, info)

	// Register with mux
	if adapter.mux != nil {
		adapter.mux.HandleFunc(pattern, handler)
	} else {
		http.HandleFunc(pattern, handler)
	}
}

// HandleFuncWithInfo registers a handler function with custom documentation.
func (adapter *NetHTTPAdapter) HandleFuncWithInfo(
	pattern string,
	handler http.HandlerFunc,
	info RouteInfo,
) {
	// Register with custom documentation
	adapter.swagger.AddRoute("GET", pattern, info)

	// Register with mux
	if adapter.mux != nil {
		adapter.mux.HandleFunc(pattern, handler)
	} else {
		http.HandleFunc(pattern, handler)
	}
}

// GET registers a GET handler with automatic documentation.
func (adapter *NetHTTPAdapter) GET(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.wrapHandler(pattern, "GET", handler)
	return adapter.swagger.GET(pattern).RegisterHandler(handler)
}

// POST registers a POST handler with automatic documentation.
func (adapter *NetHTTPAdapter) POST(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.wrapHandler(pattern, "POST", handler)
	return adapter.swagger.POST(pattern).RegisterHandler(handler)
}

// PUT registers a PUT handler with automatic documentation.
func (adapter *NetHTTPAdapter) PUT(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.wrapHandler(pattern, "PUT", handler)
	return adapter.swagger.PUT(pattern).RegisterHandler(handler)
}

// DELETE registers a DELETE handler with automatic documentation.
func (adapter *NetHTTPAdapter) DELETE(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.wrapHandler(pattern, "DELETE", handler)
	return adapter.swagger.DELETE(pattern).RegisterHandler(handler)
}

// PATCH registers a PATCH handler with automatic documentation.
func (adapter *NetHTTPAdapter) PATCH(pattern string, handler http.HandlerFunc) *RouteBuilder {
	adapter.wrapHandler(pattern, "PATCH", handler)
	return adapter.swagger.PATCH(pattern).RegisterHandler(handler)
}

// wrapHandler wraps the handler with method-specific routing.
func (adapter *NetHTTPAdapter) wrapHandler(
	pattern, method string,
	handler http.HandlerFunc,
) http.HandlerFunc {
	wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}

	// Register with mux
	if adapter.mux != nil {
		adapter.mux.HandleFunc(pattern, wrappedHandler)
	} else {
		http.HandleFunc(pattern, wrappedHandler)
	}

	return wrappedHandler
}

// ServeHTTP implements http.Handler interface.
func (adapter *NetHTTPAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if adapter.mux != nil {
		adapter.mux.ServeHTTP(w, r)
	} else {
		http.DefaultServeMux.ServeHTTP(w, r)
	}
}

// ServeDocs serves the Swagger documentation.
func (adapter *NetHTTPAdapter) ServeDocs(pattern string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		json, err := adapter.swagger.ToJSON()
		if err != nil {
			http.Error(w, "Failed to generate documentation", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}

	if adapter.mux != nil {
		adapter.mux.HandleFunc(pattern, handler)
	} else {
		http.HandleFunc(pattern, handler)
	}
}

// ServeSwaggerUI serves the Swagger UI.
func (adapter *NetHTTPAdapter) ServeSwaggerUI(pattern, docsURL string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		html := generateSwaggerUIHTML(docsURL)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}

	if adapter.mux != nil {
		adapter.mux.HandleFunc(pattern, handler)
	} else {
		http.HandleFunc(pattern, handler)
	}
}

// generateSwaggerUIHTML generates the HTML for Swagger UI.
func generateSwaggerUIHTML(docsURL string) string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui.css" />
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *, *:before, *:after {
            box-sizing: inherit;
        }
        body {
            margin:0;
            background: #fafafa;
        }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@3.52.5/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '` + docsURL + `',
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
}

