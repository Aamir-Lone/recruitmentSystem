package controllers

import (
	"RecruitmentManagementSystem/services"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func UploadResume(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := "./resumes/" + header.Filename
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	profileData, err := services.UploadResume(filePath)
	if err != nil {
		http.Error(w, "Error parsing resume", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileData)
}
