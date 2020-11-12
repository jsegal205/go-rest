package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Recipe struct {
	Slug        string `json:"slug"`
	Title       string `json:"Title"`
	Ingredients string `json:"Ingredients"`
	Directions  string `json:"Directions"`
}

var Recipes []Recipe

func allRecipes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit:: all Recipes")
	json.NewEncoder(w).Encode(Recipes)
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

	Recipes = []Recipe{
		Recipe{Slug: "cookies", Title: "really bad cookies", Ingredients: "flour, eggs, milk, sugar", Directions: "combine all the things and bake"},
		Recipe{Slug: "ham", Title: "ham", Ingredients: "ham", Directions: "just cut a slice and eat it"},
	}

	handleRequests()
}
