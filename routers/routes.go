package routers

import (
	"RecruitmentManagementSystem/controllers"
	"RecruitmentManagementSystem/utils"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Apply middleware
	r.Use(utils.JWTAuthMiddleware)

	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/uploadResume", controllers.UploadResume).Methods("POST")
	r.HandleFunc("/admin/job", controllers.CreateJob).Methods("POST")
	r.HandleFunc("/admin/job/{job_id}", controllers.GetJob).Methods("GET")
	r.HandleFunc("/admin/applicants", controllers.GetAllApplicants).Methods("GET")
}
