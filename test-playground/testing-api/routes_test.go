package main

import (
	"testing"
	"testing-api/internal/cmiddleware"
	"testing-api/internal/config"
	"testing-api/internal/handlers"

	"github.com/go-fuego/fuego"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCtx struct {
	*fuego.MockContext[any]
	tmplName string
	tmplData any
}

func newTestCtx() *testCtx {
	// Base MockContext gives you Body, QueryParam, etc.
	return &testCtx{MockContext: fuego.NewMockContextNoBody()}
}

// Override Render to capture arguments instead of panicking
func (t *testCtx) Render(
	templateToExecute string,
	data any,
	templateGlobsToOverride ...string,
) (fuego.CtxRenderer, error) {
	t.tmplName = templateToExecute
	t.tmplData = data
	// Return a no-op renderer
	return struct{ fuego.CtxRenderer }{}, nil
}

func TestHome(t *testing.T) {
	newIpMiddleware := cmiddleware.New()

	app := config.New(newIpMiddleware)
	h := handlers.New(app)

	t.Run("TestHome", func(t *testing.T) {
		ctx := newTestCtx()
		renderer, err := h.Home(ctx)
		require.NoError(t, err)
		assert.Equal(t, "home.page.html", ctx.tmplName)
		assert.Equal(t, fuego.H{"IP": "127.0.0.1"}, ctx.tmplData)
		assert.NotNil(t, renderer)
	})
}
