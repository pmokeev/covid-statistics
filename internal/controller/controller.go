package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pmokeev/covid-statistic/internal/service"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
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
