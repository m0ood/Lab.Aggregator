package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getNews(w http.ResponseWriter, r *http.Request) {
	newsData := []string{"News 1", "News 2", "News 3"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"news": newsData})
}

func main() {
	http.HandleFunc("/news", getNews)
	
	port := 8081
	fmt.Printf("Сервер новостей запущен на порту %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера новостей: %v\n", err)
	}
}
