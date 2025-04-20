package config

import "testing-api/internal/cmiddleware"

type Application struct {
	CMiddlewares *cmiddleware.MW
}

func New(m *cmiddleware.MW) *Application {
	return &Application{
		CMiddlewares: m,
	}
}
