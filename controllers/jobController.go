package controllers

import (
	"RecruitmentManagementSystem/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job services.JobRequest
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = services.CreateJob(job)
	if err != nil {
		http.Error(w, "Error creating job", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Job created successfully")
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	jobID := mux.Vars(r)["job_id"]
	job, err := services.GetJobByID(jobID)
	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(job)
}
