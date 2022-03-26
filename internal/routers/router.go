package routers

import (
	"github.com/pmokeev/covid-statistic/internal/controllers"
	"net/http"

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

	return router
}
