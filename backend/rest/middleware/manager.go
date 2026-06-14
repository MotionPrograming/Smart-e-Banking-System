package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func (m *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	if len(m.globalMiddlewares) == 0 {
		return handler
	}

	for i := len(m.globalMiddlewares) - 1; i >= 0; i-- {
		h = m.globalMiddlewares[i](h)
	}

	return h
}

func (m *Manager) With(handler http.Handler, extraMiddlewares ...Middleware) http.Handler {
	h := handler

	for i := len(extraMiddlewares) - 1; i >= 0; i-- {
		h = extraMiddlewares[i](h)
	}

	return h
}
