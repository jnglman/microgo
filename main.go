package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"microgo/container"
)

func main() {
	priceHandler := container.DependencyContainer().InjectHandler()
	router := mux.NewRouter()
	router.HandleFunc("/", priceHandler.SavePrice).Methods("POST")
	router.HandleFunc("/{id}", priceHandler.GetPrice).Methods("GET")
	err := http.ListenAndServe(":8081", router); if err != nil {
		panic(err)
	}
}
