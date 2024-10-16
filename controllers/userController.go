package controllers

import (
	"RecruitmentManagementSystem/services"
	"encoding/json"
	"net/http"
)

// UploadResume handles the resume upload API endpoint
func UploadResume(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "userID not found", http.StatusBadRequest)
		return
	}

	//userID := r.Context().Value("userID").(string) // Assuming you set userID in the context during authentication

	// Parse the form to get file
	err := r.ParseMultipartForm(10 << 20) // Limit to 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from form input
	file, fileHeader, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Unable to get resume file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Call the service to upload the resume
	err = services.UploadResume(userID, file, fileHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Resume uploaded successfully"))
}

// GetAllApplicants handles the API endpoint to retrieve all applicants
func GetAllApplicants(w http.ResponseWriter, r *http.Request) {
	applicants, err := services.GetAllApplicants()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return applicants as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(applicants)
}
