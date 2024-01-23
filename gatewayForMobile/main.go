package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var newsServiceURL = "http://localhost:8081/news"
var weatherServiceURL = "http://localhost:8082/weather"

type weatherData struct {
	Weather      weatherInfo `json:"weather"`
}

type weatherInfo struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
}

type newsData struct {
	News []string `json:"news"`
}

var newsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "News",
	Fields: graphql.Fields{
		"news": &graphql.Field{
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				response, err := http.Get(newsServiceURL)
				if err != nil {
					return nil, err
				}
				defer response.Body.Close()

				var data newsData
				if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
					return nil, err
				}

				return data.News, nil
			},
		},
	},
})

var weatherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Weather",
	Fields: graphql.Fields{
		"city": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				response, err := http.Get(weatherServiceURL)
				if err != nil {
					return nil, err
				}
				defer response.Body.Close()

				var data weatherData
				if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
					return nil, err
				}

				return data.Weather.City, nil
			},
		},
		"temperature": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				response, err := http.Get(weatherServiceURL)
				if err != nil {
					return nil, err
				}
				defer response.Body.Close()

				var data weatherData
				if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
					return nil, err
				}

				return data.Weather.Temperature, nil
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"news": &graphql.Field{
			Type: newsType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				response, err := http.Get(newsServiceURL)
				if err != nil {
					return nil, err
				}
				defer response.Body.Close()
				fmt.Println("News Service Response:", response.Status)
				var data newsData
				if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
					return nil, err
				}
				return data.News, nil
			},
		},
		"weather": &graphql.Field{
			Type: weatherType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				response, err := http.Get(weatherServiceURL)
				if err != nil {
					return nil, err
				}
				defer response.Body.Close()
				fmt.Println("Weather Service Response:", response.Status)
				var data weatherData
				if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
					return nil, err
				}
				return data.Weather, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	h.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/graphql", graphqlHandler)

	port := 8080
	fmt.Printf("GraphQL Gateway is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting GraphQL Gateway: %v\n", err)
	}
}
