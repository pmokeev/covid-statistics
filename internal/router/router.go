package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pmokeev/covid-statistic/internal/controller"
)

type Router struct {
	controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{controller: controller}
}

func (r *Router) InitRouter() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/", r.controller.GetStatistic)
	}

	return router
}
