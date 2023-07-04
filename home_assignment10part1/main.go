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
	url, err := url.Parse(weatherBaseUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	city := mux.Vars(r)["city"]

	q := url.Query()
	q.Set("key", apiKey)
	q.Set("q", city)
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())

	defer resp.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type WeatherResponse struct {
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
		Current struct {
			TemperatureC float64 `json:"temp_c"`
			TemperatureF float64 `json:"temp_f"`
			Condition    struct {
				Text string `json:"text"`
			} `json:"condition"`
			Humidity int `json:"humidity"`
		} `json:"current"`
	}

	var weatherResp WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&weatherResp)

	if err != nil {
		log.Default().Printf("Failed to parse exchange response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	weatherData := struct {
		TemperatureC float64 `json:"temp_c"`
		TemperatureF float64 `json:"temp_f"`
		Condition    struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity int `json:"humidity"`
	}{
		TemperatureC: weatherResp.Current.TemperatureC,
		TemperatureF: weatherResp.Current.TemperatureF,
		Condition:    weatherResp.Current.Condition,
		Humidity:     weatherResp.Current.Humidity,
	}

	err = json.NewEncoder(w).Encode(weatherData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
