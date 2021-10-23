package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	OrderID     string  `json:"orderID"`
	Description string  `json:"description"`
	Driver      string  `json:"driver"`
	Revenue     float64 `json:"revenue"`
	Cost        float64 `json:"cost"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint: Homepage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT", "OPTIONS")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("Endpoint Hit: Return all Articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.OrderID == key {
			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Fprintf(w, "")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]
	for index, article := range Articles {
		if article.OrderID == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	key := vars["id"]
	var updatedEvent Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)

	for i, article := range Articles {
		if article.OrderID == key {
			article.OrderID = updatedEvent.OrderID
			article.Description = updatedEvent.Description
			article.Driver = updatedEvent.Driver
			article.Revenue = updatedEvent.Revenue
			article.Cost = updatedEvent.Cost

			Articles[i] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}

func main() {
	Articles = []Article{
		Article{
			OrderID:     "id1",
			Description: "Construction Materials",
			Driver:      "SteveWilliams",
			Revenue:     4200.00,
			Cost:        100.00,
		},

		Article{
			OrderID:     "id2",
			Description: "Construction Materials",
			Driver:      "SteveWilliams",
			Revenue:     3948.45,
			Cost:        71.38,
		},

		Article{
			OrderID:     "id3",
			Description: "Wood and Lumber",
			Driver:      "SteveWilliams",
			Revenue:     1950.52,
			Cost:        263.88,
		},

		Article{
			OrderID:     "id4",
			Description: "Wood and Lumber",
			Driver:      "SteveWilliams",
			Revenue:     4991.45,
			Cost:        116.98,
		},

		Article{
			OrderID:     "id5",
			Description: "Meat",
			Driver:      "ChrisHorton",
			Revenue:     6739.72,
			Cost:        279.17,
		},

		Article{
			OrderID:     "id6",
			Description: "Meat",
			Driver:      "ChrisHorton",
			Revenue:     3618.08,
			Cost:        537.91,
		},

		Article{
			OrderID:     "id7",
			Description: "Fresh Produce",
			Driver:      "ChrisHorton",
			Revenue:     5345.91,
			Cost:        420.69,
		},

		Article{
			OrderID:     "id8",
			Description: "Farm Supplies",
			Driver:      "ChrisHorton",
			Revenue:     7429.78,
			Cost:        171.13,
		},

		Article{
			OrderID:     "id9",
			Description: "Cheetos",
			Driver:      "ChrisHorton",
			Revenue:     7231.98,
			Cost:        310.38,
		},

		Article{
			OrderID:     "id10",
			Description: "Rose Rocket Swag",
			Driver:      "AlexNovak",
			Revenue:     5404.24,
			Cost:        350.79,
		},
	}
	handleRequests()
}
