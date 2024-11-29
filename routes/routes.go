package routes

import (
	"encoding/json"
	"net/http"
	"weather-api/config"
	"weather-api/models"

	"github.com/gorilla/mux"
)

// GetWeather
// @Summary Get weather data by city
// @Description Get weather data for a specific city from WeatherStack API. This endpoint fetches the current weather information, including temperature and humidity, for the provided city.
// @Tags weather
// @Accept  json
// @Produce  json
// @Param city path string true "City name" // City name parameter in the URL path
// @Success 200 {object} models.Weather "Weather data successfully retrieved" // Response structure when successful
// @Failure 404 {string} string "City not found" // When the city is not found or invalid
// @Failure 500 {string} string "Internal Server Error" // In case of a server error
// @Router /weather/{city} [get] // Path and HTTP method for this route
func GetWeather(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		city := mux.Vars(r)["city"]
		weather, err := models.GetWeatherFromWeatherStack(city, apiKey)
		if err != nil {
			http.Error(w, "City not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(weather); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

// CreateWeather
// @Summary Save weather data to MongoDB
// @Description Save the provided weather data to the MongoDB database. This endpoint stores the weather data for a specific city in the database for later use and retrieval.
// @Tags weather
// @Accept  json
// @Produce  json
// @Param weather body models.Weather true "Weather data" // The weather data object to be stored in the database
// @Success 201 {string} string "Weather data successfully created" // Response message when weather data is successfully saved
// @Failure 400 {string} string "Invalid input format" // When the input data is malformed or invalid
// @Failure 500 {string} string "Failed to save weather data" // If there is an error while saving the data
// @Router /weather [post] // Path and HTTP method for this route
func CreateWeather(w http.ResponseWriter, r *http.Request) {
	var weather models.Weather
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	err = models.SaveWeather(config.Client, &weather)
	if err != nil {
		http.Error(w, "Failed to save weather data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
