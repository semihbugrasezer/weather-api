package models

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"time"
)

type Weather struct {
	City        string    `json:"city" bson:"city"`
	Temperature float64   `json:"temperature" bson:"temperature"`
	Humidity    float64   `json:"humidity" bson:"humidity"`
	Date        string    `json:"date" bson:"date"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

type WeatherStackResponse struct {
	Request struct {
		Query string `json:"query"`
	} `json:"request"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temperature float64 `json:"temperature"`
		Humidity    float64 `json:"humidity"`
	} `json:"current"`
}

func GetWeatherFromWeatherStack(city string, apiKey string) (Weather, error) {
	var apiResponse WeatherStackResponse
	var weather Weather

	url := fmt.Sprintf("https://api.weatherstack.com/current?access_key=%s&query=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return weather, fmt.Errorf("HTTP isteği başarısız: %v", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("Yanıt kapatılırken hata oluştu:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("API'den veri alınamadı: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return weather, fmt.Errorf("yanıt JSON'u çözümlenemedi: %v", err)
	}

	weather = Weather{
		City:        apiResponse.Location.Name,
		Temperature: apiResponse.Current.Temperature,
		Humidity:    apiResponse.Current.Humidity,
		Date:        time.Now().Format("2006-01-02"),
		CreatedAt:   time.Now(),
	}

	return weather, nil
}

func CreateTTLIndex(client *mongo.Client) error {
	collection := client.Database("weatherDB").Collection("weather")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "createdAt", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(86400), // 24 saatlik TTL
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return fmt.Errorf("TTL indeksi oluşturulamadı: %v", err)
	}
	return nil
}

func SaveWeather(client *mongo.Client, weather *Weather) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database("weatherDB").Collection("weather")

	filter := bson.M{"city": weather.City, "date": weather.Date}
	existing := collection.FindOne(ctx, filter)
	if existing.Err() == nil {
		log.Printf("Bu şehir (%s) için veri zaten mevcut.", weather.City)
		return nil
	}

	_, err := collection.InsertOne(ctx, weather)
	if err != nil {
		return fmt.Errorf("veri kaydedilemedi: %v", err)
	}
	log.Printf("%s şehri için veri başarıyla kaydedildi.", weather.City)
	return nil
}
