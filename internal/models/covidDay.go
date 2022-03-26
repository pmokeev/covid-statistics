package models

import "time"

type CovidDay struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
}
