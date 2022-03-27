package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pmokeev/covid-statistic/internal/models"
)

type Service struct {
	client *redis.Client
}

func NewService() *Service {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &Service{client: client}
}

func (s *Service) GetStatistic(ctx context.Context, country string) ([]models.CovidDay, error) {
	queryString := "https://api.covid19api.com/country/%s/status/confirmed?from=%s&to=%s"

	currentTime := time.Now()
	endDate := currentTime.Format("2006-01-02")
	startDate := currentTime.AddDate(0, -1, 0).Format("2006-01-02")

	queryString = fmt.Sprintf(queryString, country, startDate, endDate)
	value, err := s.client.Get(ctx, queryString).Result()
	if err == redis.Nil {
		response, err := http.Get(queryString)
		if err != nil {
			return nil, err
		}

		data := make([]models.CovidDay, 0)
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return nil, err
		}

		refactoredData := make([]models.CovidDay, 0)
		currentDay := data[0].Date
		counterDay := 0
		for _, value := range data {
			if value.Date == currentDay {
				counterDay += value.Cases
			} else {
				refactoredData = append(refactoredData, models.CovidDay{Cases: counterDay,
					Date: currentDay})
				counterDay = value.Cases
				currentDay = value.Date
			}
		}

		byteData, err := json.Marshal(refactoredData)
		if err != nil {
			return nil, err
		}

		err = s.client.Set(ctx, queryString, bytes.NewBuffer(byteData).Bytes(), time.Hour*2).Err()
		if err != nil {
			return nil, err
		}

		return refactoredData, nil
	} else if err != nil {
		return nil, err
	} else {
		data := make([]models.CovidDay, 0)

		err := json.Unmarshal(bytes.NewBufferString(value).Bytes(), &data)
		if err != nil {
			return nil, err
		}

		return data, err
	}
}
