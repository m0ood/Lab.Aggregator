package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func aggregateDataHandler(w http.ResponseWriter, r *http.Request) {
	newsData := fetchDataFromService("http://localhost:8081/news")
	weatherData := fetchDataFromService("http://localhost:8082/weather")
	aggregatedData := map[string]interface{}{"news": newsData["news"], "weather": weatherData["weather"]}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aggregatedData)
}

func fetchDataFromService(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{"error": "Ошибка получения данных от сервиса"}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"error": "Ошибка чтения данных от сервиса"}
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return map[string]interface{}{"error": "Ошибка обработки данных от сервиса"}
	}

	return data
}

func main() {
	http.HandleFunc("/aggregate", aggregateDataHandler)

	port := 8080
	fmt.Printf("Сервер агрегатора запущен на порту %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера агрегатора: %v\n", err)
	}
}
