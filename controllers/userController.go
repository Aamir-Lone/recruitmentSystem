package controllers

import (
	"RecruitmentManagementSystem/services"
	"io"
	"net/http"
	"os"
)

func UploadResumeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // Limit the file size to 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, fileHeader, err := r.FormFile("resume") // Get both file and its header
	if err != nil {
		http.Error(w, "Unable to retrieve file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file to a temporary location (or process as needed)
	tempFilePath := "uploads/" + fileHeader.Filename // Use fileHeader.Filename to get the file name

	// Create the uploads directory if it doesn't exist
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		http.Error(w, "Unable to create uploads directory", http.StatusInternalServerError)
		return
	}

	out, err := os.Create(tempFilePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// Extract userID from context (make sure you have AuthMiddleware set up correctly)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Call the service to process the resume
	err = services.UploadResume(tempFilePath, userID) // Pass the userID correctly
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Resume uploaded successfully"))
}
