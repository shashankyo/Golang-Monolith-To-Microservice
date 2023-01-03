package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/middleware"
	"github.com/gofiber/fiber/middleware"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return r
}
