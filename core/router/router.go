package routes

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/maurana/nuswantara/core/config"
	"github.com/maurana/nuswantara/core/security/jwt"
	chimiddleware "github.com/go-chi/chi/middleware"
)


func NuswantaraRouter() (*chi.Mux, chi.Router) {
	router := chi.NewRouter()

	router.Use(httprate.LimitByIP(
		config.Cfg().HttpRateLimitRequest,
		config.Cfg().HttpRateLimitTime,
	))

	router.Use(cors.AllowAll().Handler)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
    router.Options("/*", func(w http.ResponseWriter, r *http.Request) {})

    return router, chi.Router
}
