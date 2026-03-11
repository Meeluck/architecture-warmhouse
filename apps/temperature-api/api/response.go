package api

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"
)

type TemperatureResponse struct {
	Location    string  `json:"location"`
	SensorId    string  `json:"sensorId"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Timestamp   string  `json:"timestamp"`
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorId := r.URL.Query().Get("sensorId")
	if location == "" {
		switch sensorId {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	if sensorId == "" {
		switch location {
		case "Living Room":
			sensorId = "1"
		case "Bedroom":
			sensorId = "2"
		case "Kitchen":
			sensorId = "3"
		default:
			sensorId = "0"
		}
	}
	temperature := randomTemperature(location)

	resp := TemperatureResponse{
		Location:    location,
		SensorId:    sensorId,
		Temperature: temperature,
		Unit:        "C",
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func randomTemperature(location string) float64 {
	var min, max float64

	switch location {
	case "Living Room":
		min, max = 18.0, 25.0
	case "Bedroom":
		min, max = 16.0, 23.0
	case "Kitchen":
		min, max = 20.0, 30.0
	default:
		min, max = 15.0, 35.0
	}

	value := min + rand.Float64()*(max-min)
	result, _ := strconv.ParseFloat(strconv.FormatFloat(value, 'f', 1, 64), 64)
	return result
}
