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
	country, status := context.GetQuery("country")
	if !status {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	statistic, err := c.service.GetStatistic(country)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	context.JSON(http.StatusOK, statistic)
}
