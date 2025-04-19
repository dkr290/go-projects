package main

import (
	"testing"

	"github.com/go-fuego/fuego"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	ctx := fuego.NewMockContextNoBody()
	resp, err := Home(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", resp)
}

func Home(c fuego.ContextNoBody) (string, error) {
	return "Hello, World!", nil
}
