Here is a detailed README.md file for your project:

# Weather API

The Weather API is a RESTful API that provides weather data for specific cities and saves the data to a MongoDB database.

## Features

- **Fetch Weather Data:** Retrieve weather data for a specific city from the WeatherStack API, including temperature and humidity.
- **Save Weather Data:** Save weather data for a city in a MongoDB database for future reference.
- **Automatic Data Expiry:** Implements a TTL index to automatically remove data older than 24 hours from the database.
- **Rate Limiting:** Includes rate limiting to prevent excessive API usage.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/weather-api.git
   cd weather-api

	2.	Install dependencies:

go mod tidy


	3.	Create a .env file in the root directory and add the following variables:

MONGODB_URI=your_mongodb_connection_string
WEATHERSTACK_API_KEY=your_weatherstack_api_key


	4.	Run the API:

go run main.go



API Endpoints

1. Get Weather Data by City

	•	URL: /weather/{city}
	•	Method: GET
	•	Description: Fetch weather data for a specific city from the WeatherStack API.
	•	Parameters:
	•	city (path): Name of the city to fetch weather data for.
	•	Responses:
	•	200 OK: Returns the weather data for the city.
	•	404 Not Found: City not found.
	•	500 Internal Server Error: Server error.

2. Save Weather Data

	•	URL: /weather
	•	Method: POST
	•	Description: Save weather data to the MongoDB database.
	•	Request Body:

{
  "city": "City Name",
  "temperature": 25.5,
  "humidity": 60,
  "date": "2024-01-01",
  "createdAt": "2024-01-01T12:00:00Z"
}


	•	Responses:
	•	201 Created: Weather data successfully saved.
	•	400 Bad Request: Invalid input format.
	•	500 Internal Server Error: Failed to save weather data.

Project Structure

	•	config/: Contains the MongoDB connection logic.
	•	models/: Includes data models and functions for interacting with MongoDB and WeatherStack API.
	•	routes/: Defines HTTP handlers for the API endpoints.
	•	main.go: The entry point of the application.
	•	docs/: Swagger documentation files.

Swagger Documentation

Swagger UI is available at /swagger/. It provides detailed API documentation and allows testing the endpoints interactively.

Example Usage

Fetch Weather Data

curl -X GET http://localhost:8080/weather/London

Save Weather Data

curl -X POST http://localhost:8080/weather \
-H "Content-Type: application/json" \
-d '{
  "city": "London",
  "temperature": 25.5,
  "humidity": 60,
  "date": "2024-01-01",
  "createdAt": "2024-01-01T12:00:00Z"
}'

Dependencies

	•	Go Modules: Dependency management.
	•	gorilla/mux: HTTP router.
	•	swaggo/swag: Swagger documentation.
	•	WeatherStack API: For fetching weather data.
	•	MongoDB: Database for storing weather data.

License

This project is licensed under the MIT License. See the LICENSE file for details.
