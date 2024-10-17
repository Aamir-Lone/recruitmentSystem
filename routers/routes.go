package routers

import (
	"RecruitmentManagementSystem/controllers"
	"RecruitmentManagementSystem/utils"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/uploadResume", utils.AuthMiddleware(controllers.UploadResumeHandler)).Methods("POST") // Use the updated handler
	r.HandleFunc("/admin/job", utils.AdminAuthMiddleware(controllers.CreateJob)).Methods("POST")

	return r
}
