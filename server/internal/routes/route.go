package routes

import (
	"github.com/SaikatDeb12/ecommerce/internal/handler"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/health", handler.CheckHealth)
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", handler.SignUp)
			r.Post("/signup", handler.SignIn)
		})
	})

	return router
}
