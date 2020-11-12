package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Recipe struct {
	Title       string `json:"Title"`
	Ingredients string `json:"Ingredients"`
	Directions  string `json:"Directions"`
}

type Recipes []Recipe

func allRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := Recipes{
		Recipe{Title: "really bad cookies", Ingredients: "flour, eggs, milk, sugar", Directions: "combine all the things and bake"},
		Recipe{Title: "ham", Ingredients: "ham", Directions: "just cut a slice and eat it"},
	}

	fmt.Println("endpoint hit:: all Recipes")
	json.NewEncoder(w).Encode(recipes)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello jimbo jellybeans")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/recipes", allRecipes)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO Rest API")
	handleRequests()
}
