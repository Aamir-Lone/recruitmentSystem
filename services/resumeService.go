package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"RecruitmentManagementSystem/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	err = storeExtractedData(resumeData)
	if err != nil {
		return err
	}

	return nil
}

// extractResumeData calls the third-party API to extract resume data
// func extractResumeData(resumePath string) (models.Profile, error) {
// 	file, err := os.Open(resumePath)
// 	if err != nil {
// 		return models.Profile{}, fmt.Errorf("could not open resume file: %v", err)
// 	}
// 	defer file.Close()

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", file)
// 	if err != nil {
// 		return models.Profile{}, fmt.Errorf("could not create request: %v", err)
// 	}

// 	req.Header.Set("Content-Type", "application/octet-stream")
// 	req.Header.Set("apikey", "0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return models.Profile{}, fmt.Errorf("could not send request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return models.Profile{}, fmt.Errorf("could not read response body: %v", err)
// 	}

// 	var resumeData models.Profile
// 	if err := json.Unmarshal(body, &resumeData); err != nil {
// 		return models.Profile{}, fmt.Errorf("could not unmarshal response: %v", err)
// 	}

//		return resumeData, nil
//	}

// storeExtractedData saves the extracted resume data in the database

func storeExtractedData(data models.Profile) error {
	// Log the extracted profile data
	fmt.Println("Extracted Profile Data:", data)

	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Aamirlone:Aamirlone@cluster0.1zwiz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		return fmt.Errorf("could not connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Get the user collection
	collection := client.Database("recruitment_management").Collection("users")

	// Create a new User object with the profile data
	newUser := models.User{
		Email:   data.Phone, // Use an appropriate field for the email
		Profile: data,       // Assign the profile data from the resume
	}

	// Insert the new user document into the collection
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return fmt.Errorf("could not insert new user: %v", err)
	}

	// Log the ID of the inserted document
	fmt.Println("Inserted document ID:", result.InsertedID)

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

	// Log the status and response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Profile{}, fmt.Errorf("could not read response body: %v", err)
	}

	// Log the response body for debugging
	fmt.Println("API Response:", string(body))

	var resumeData models.Profile
	if err := json.Unmarshal(body, &resumeData); err != nil {
		return models.Profile{}, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return resumeData, nil
}
