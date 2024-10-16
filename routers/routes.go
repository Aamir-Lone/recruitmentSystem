package routers

import (
	"RecruitmentManagementSystem/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	r.HandleFunc("/uploadResume", controllers.UploadResume).Methods("POST")
	r.HandleFunc("/admin/job", controllers.CreateJob).Methods("POST")
	r.HandleFunc("/admin/job/{job_id}", controllers.GetJob).Methods("GET")
	r.HandleFunc("/admin/applicants", controllers.GetAllApplicants).Methods("GET")
}
