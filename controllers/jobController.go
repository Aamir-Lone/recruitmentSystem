package controllers

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/services"
	"encoding/json"
	"net/http"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	json.NewDecoder(r.Body).Decode(&job)
	createdJob, err := services.CreateJob(job)
	if err != nil {
		http.Error(w, "Unable to create job", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdJob)
}
