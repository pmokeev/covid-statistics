package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pmokeev/covid-statistic/internal/services"
	"net/http"
)

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetStatistic(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	statistic, err := c.service.GetStatistic(r.Context(), country)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(statistic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
