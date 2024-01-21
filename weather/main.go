package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	weatherData := map[string]string{"city": "Moscow", "temperature": "-15°C"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"weather": weatherData})
}

func main() {
	http.HandleFunc("/weather", getWeather)

	port := 8082
	fmt.Printf("Сервер погоды запущен на порту %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера погоды: %v\n", err)
	}
}
