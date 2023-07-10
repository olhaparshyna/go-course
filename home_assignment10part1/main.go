package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

//1. Який отримує запити на адресі "/weather" методом GET та повертає погодні дані для заданого міста.
//Додаткові вимоги:
//◦ використовувати сторонній API, який надає погодні дані;
//◦ запит до API має включати параметр, який вказує місто, для якого повертаються погодні дані;
//◦ результат повертати у вигляді JSON об'єкту, де ключі — це погодні параметри (температура, вітер, вологість тощо), а значення — це відповідні значення.

func main() {
	r := mux.NewRouter()

	weatherRes := &weatherResource{}

	r.HandleFunc("/weather/{city}", weatherRes.getWeatherByCity).
		Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

const (
	weatherBaseUrl = "https://api.weatherapi.com/v1/current.json"
	apiKey         = "a3dc4a2951414b6cb57175549230307"
)

type weatherResource struct {
}

func (wr *weatherResource) getWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]

	weatherData, err := getWeatherData(city)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(weatherData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type WeatherData struct {
	TemperatureC float64 `json:"temp_c"`
	TemperatureF float64 `json:"temp_f"`
	Condition    struct {
		Text string `json:"text"`
	} `json:"condition"`
	Humidity int `json:"humidity"`
}

type Location struct {
	Name string `json:"name"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TemperatureC float64   `json:"temp_c"`
	TemperatureF float64   `json:"temp_f"`
	Condition    Condition `json:"condition"`
	Humidity     int       `json:"humidity"`
}

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func getWeatherData(city string) (*WeatherData, error) {
	url, err := url.Parse(weatherBaseUrl)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("key", apiKey)
	q.Set("q", city)
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var weatherResp WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&weatherResp)

	if err != nil {
		return nil, err
	}

	weatherData := &WeatherData{
		TemperatureC: weatherResp.Current.TemperatureC,
		TemperatureF: weatherResp.Current.TemperatureF,
		Condition:    weatherResp.Current.Condition,
		Humidity:     weatherResp.Current.Humidity,
	}

	return weatherData, nil
}
