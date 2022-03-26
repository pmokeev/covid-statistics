package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmokeev/covid-statistic/internal/service"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetStatistic(context *gin.Context) {
	context.AbortWithStatus(http.StatusBadRequest)
}
