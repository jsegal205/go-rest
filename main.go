package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func singleRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	fmt.Println("endpoint hit:: single Recipe for " + slug)

	for _, recipe := range Recipes {
		if recipe.Slug == slug {
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}

	fmt.Fprintf(w, "404 Recipe not found for "+slug)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("endpoint hit:: create Recipe with " + string(reqBody))

	// TODO ::
	// add logic to disallow slug duplication
	// add required field logic
	// add auth logic

	var newRecipe Recipe
	json.Unmarshal(reqBody, &newRecipe)

	Recipes = append(Recipes, newRecipe)

	json.NewEncoder(w).Encode(newRecipe)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	fmt.Println("endpoint hit:: delete Recipe for " + slug)

	// TODO ::
	// add auth logic

	for index, recipe := range Recipes {
		if recipe.Slug == slug {
			// I understand this takes all the recipes up until the index and appends
			// with the all elements after the index. I believe the `...` means all the
			// rest of the elements
			Recipes = append(Recipes[:index], Recipes[index+1:]...)
			return
		}
	}
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	fmt.Println("endpoint hit:: update Recipe for " + slug)

	// TODO ::
	// add logic to disallow slug being changed
	// add logic for required fields being there if updating to nil
	// add auth logic

	for index, recipe := range Recipes {
		if recipe.Slug == slug {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var updatedRecipe Recipe
			json.Unmarshal(reqBody, &updatedRecipe)
			Recipes[index] = updatedRecipe
			return
		}
	}

	fmt.Fprintf(w, "404 Recipe not found for "+slug)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello jimbo jellybeans")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/recipes", allRecipes)
	myRouter.HandleFunc("/recipe", createRecipe).Methods("POST")
	myRouter.HandleFunc("/recipe/{slug}", deleteRecipe).Methods("DELETE")
	myRouter.HandleFunc("/recipe/{slug}", updateRecipe).Methods("PUT")
	myRouter.HandleFunc("/recipe/{slug}", singleRecipe)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO Rest API")

	// This global `Recipes` mimics data from a database without all the hassle of
	// setting up a database
	Recipes = []Recipe{
		Recipe{Slug: "cookies", Title: "really bad cookies", Ingredients: "flour, eggs, milk, sugar", Directions: "combine all the things and bake"},
		Recipe{Slug: "ham", Title: "ham", Ingredients: "ham", Directions: "just cut a slice and eat it"},
	}

	handleRequests()
}
