package main

import (
	"RecruitmentManagementSystem/routers"
	"RecruitmentManagementSystem/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access sensitive data using os.Getenv
	//mongoDBURI := os.Getenv("MONGODB_URI")
	//apiKey := os.Getenv("API_KEY")

	utils.ConnectDB()

	r := routers.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
