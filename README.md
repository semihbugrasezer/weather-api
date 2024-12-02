# 🌦️ Weather API

A simple REST API for fetching and storing weather data using Go, Gorilla Mux, and MongoDB. This project demonstrates how to interact with an external API (WeatherStack), store data in MongoDB, and serve weather-related endpoints.

## Features

- **Get Weather by City**: Fetch current weather (temperature and humidity) for a specific city from the WeatherStack API.
- **Save Weather Data**: Store weather data for later retrieval using MongoDB.
- **Automatic Cleanup**: Old weather data is automatically deleted using MongoDB TTL indexing.
- **Rate Limiting**: Prevent excessive requests with built-in rate limiting.
- **Swagger Integration**: View detailed, interactive API documentation.

## Technologies Used

- **Go** (v1.23.3)
- **Gorilla Mux** - A robust HTTP router and URL matcher for building RESTful APIs.
- **MongoDB** - A NoSQL database for efficient data storage.
- **WeatherStack API** - An external weather data provider.
- **Swagger** - For API documentation generation.

## Getting Started

### Prerequisites

- Install Go (version 1.23.3 or later).
- Install MongoDB and ensure it's running locally on port `27017`.
- Obtain a WeatherStack API key from [WeatherStack](https://weatherstack.com/).

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/weather-api.git
    cd weather-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up environment variables:
   Create a `.env` file in the project root with the following:
    ```plaintext
    MONGODB_URI=mongodb://localhost:27017/weatherdb
    WEATHERSTACK_API_KEY=your_weatherstack_api_key
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

   The server will be running on `http://localhost:8080`.

## API Endpoints

### `GET /weather/{city}`
- **Description**: Retrieves current weather data for a specified city.
- **Path Parameter**:
    - `city`: Name of the city (required).
- **Response**:
    ```json
    {
        "city": "London",
        "temperature": 15.5,
        "humidity": 80,
        "date": "2024-01-01",
        "createdAt": "2024-01-01T12:00:00Z"
    }
    ```

### `POST /weather`
- **Description**: Stores weather data in MongoDB.
- **Request Body**:
    ```json
    {
        "city": "London",
        "temperature": 15.5,
        "humidity": 80,
        "date": "2024-01-01",
        "createdAt": "2024-01-01T12:00:00Z"
    }
    ```
- **Response**:
    ```json
    {
        "message": "Weather data successfully created"
    }
    ```

## Swagger Documentation

- Access the interactive Swagger documentation at:

http://localhost:8080/swagger/

## Contribution

Contributions are welcome! Fork this repository, make your changes, and open a pull request. Whether it's fixing bugs, improving documentation, or adding new features, every contribution counts.

## License

This project is open-source and available under the [MIT License](LICENSE).
