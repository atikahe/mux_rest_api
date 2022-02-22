package main

import (
	"log"
	"mux_rest_api/configs"
	"mux_rest_api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Run db
	configs.ConnectDB()

	// Route endpoints
	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "application/json")

	// 	json.NewEncoder(rw).Encode(map[string]string{"data": "REST API with Mux"})
	// }).Methods("GET")

	// Routes
	routes.UserRoute(router)

	// Listen in port
	log.Fatal(http.ListenAndServe(":6000", router))
}
