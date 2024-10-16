package main

import (
	"log"
	"net/http"

	"RecruitmentManagementSystem/routers"
	"RecruitmentManagementSystem/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	utils.ConnectDB()
	r := mux.NewRouter()
	defer utils.DisconnectDB() // Ensure disconnection when the program exits

	// Set up routes and start the server
	routers.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
