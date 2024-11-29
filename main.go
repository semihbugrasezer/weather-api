package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/juju/ratelimit"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"weather-api/config"
	_ "weather-api/docs"
	"weather-api/models"
	"weather-api/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	apiKey := os.Getenv("WEATHERSTACK_API_KEY")

	if mongoURI == "" || apiKey == "" {
		log.Fatal("MONGODB_URI veya WEATHERSTACK_API_KEY tanımlı değil.")
	}

	config.ConnectMongoDB(mongoURI)
	defer config.CloseMongoDB()

	err := models.CreateTTLIndex(config.Client)
	if err != nil {
		log.Fatalf("TTL indeksi oluşturulamadı: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", routes.GetWeather(apiKey)).Methods("GET")
	r.HandleFunc("/weather", routes.CreateWeather).Methods("POST")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	bucket := ratelimit.NewBucketWithRate(1, 5)
	r.Use(RateLimiter(bucket))

	fmt.Println("API calisiyor...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func RateLimiter(bucket *ratelimit.Bucket) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if bucket.TakeAvailable(1) == 0 {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
