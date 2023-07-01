package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "Corolla", 10000},
	{2, "Toyota", "Camy", 20000},
	{3, "Honda", "Civic", 15000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", router))
}
