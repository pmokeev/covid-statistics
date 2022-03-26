package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmokeev/covid-statistic/internal/controller"
)

type Router struct {
	controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{controller: controller}
}

func (r *Router) InitRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/", r.controller.GetStatistic).Methods("GET")

	return router
}
