package main

import (
	"context"
	"net/http/httptest"
	"testing"
	"testing-api/internal/cmiddleware"
	"testing-api/internal/config"
	"testing-api/internal/handlers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-fuego/fuego"
	"github.com/stretchr/testify/require"
)

type contextKey string

const ipContextKey contextKey = "ip"

type MW struct{}

// Simulated version of GetIpFromContext to match usage in your app.
func (m *MW) GetIpFromContext(ctx context.Context) string {
	ip, ok := ctx.Value(ipContextKey).(string)
	if !ok || ip == "" {
		return "unknown"
	}
	return ip
}

func TestNewServer(t *testing.T) {
	newIpMiddleware := cmiddleware.New()

	app := config.New(newIpMiddleware)
	s := fuego.NewServer(

		fuego.WithGlobalMiddlewares(middleware.Recoverer, app.CMiddlewares.AddIpToContext),
	)
	h := handlers.New(app)

	t.Run("can register controller", func(t *testing.T) {
		fuego.Get(s, "/", h.Home)
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		ctx := context.WithValue(req.Context(), ipContextKey, "127.0.0.1")
		req = req.WithContext(ctx)
		recorder := httptest.NewRecorder()

		s.Mux.ServeHTTP(recorder, req)

		require.Equal(t, 200, recorder.Code)
	})
}
