package services

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"RecruitmentManagementSystem/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadResume uploads a user's resume and extracts data
func UploadResume(userID string, file multipart.File, fileHeader *multipart.FileHeader) error {
	// Create a directory to store resumes if it doesn't exist
	resumeDir := "./resumes"
	if err := os.MkdirAll(resumeDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create resumes directory: %v", err)
	}

	// Save the resume file
	resumePath := filepath.Join(resumeDir, fileHeader.Filename)
	out, err := os.Create(resumePath)
	if err != nil {
		return fmt.Errorf("could not create resume file: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return fmt.Errorf("could not save resume file: %v", err)
	}

	// Call third-party API to extract resume details
	resumeData, err := extractResumeData(resumePath)
	if err != nil {
		return err
	}

	// Store extracted data in the database
	err = storeExtractedData(userID, resumeData)
	if err != nil {
		return err
	}

	return nil
}

// extractResumeData calls the third-party API to extract resume data
func extractResumeData(resumePath string) (models.Profile, error) {
	file, err := os.Open(resumePath)
	if err != nil {
		return models.Profile{}, fmt.Errorf("could not open resume file: %v", err)
	}
	defer file.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", file)
	if err != nil {
		return models.Profile{}, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", "0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS")

	resp, err := client.Do(req)
	if err != nil {
		return models.Profile{}, fmt.Errorf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Profile{}, fmt.Errorf("could not read response body: %v", err)
	}

	var resumeData models.Profile
	if err := json.Unmarshal(body, &resumeData); err != nil {
		return models.Profile{}, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return resumeData, nil
}

// storeExtractedData saves the extracted resume data in the database
func storeExtractedData(userID string, data models.Profile) error {
	// Convert userID to ObjectID
	/*objID*/
	_, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	// Here you would typically use your database package to update the user's profile with the extracted data
	// Example: Update user profile in MongoDB
	// ...
	return nil
}
