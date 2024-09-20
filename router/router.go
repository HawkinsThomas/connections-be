package router

import (
    "connections/handlers"
		"github.com/go-chi/chi/v5"
		"github.com/go-chi/chi/v5/middleware"
)

func InitializeRouter() *chi.Mux {
    r := chi.NewRouter()

    // Register routes
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
    r.Get("/health", handlers.HealthHandler)
    r.Get("/generateAnswers", handlers.GenerateAnswersHandler)

    return r
}