package main

import (
	"RecruitmentManagementSystem/routers"
	"RecruitmentManagementSystem/utils"
	"log"
	"net/http"
)

func main() {
	utils.ConnectDB()

	r := routers.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
