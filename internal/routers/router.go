package routers

import (
	"net/http"

	"github.com/pmokeev/covid-statistic/internal/controllers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

type Router struct {
	controller *controllers.Controller
}

func NewRouter(controller *controllers.Controller) *Router {
	return &Router{controller: controller}
}

func (r *Router) InitRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/", r.controller.GetStatistic).Methods("GET")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{
			http.MethodGet,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	return corsOpts.Handler(router)
}
