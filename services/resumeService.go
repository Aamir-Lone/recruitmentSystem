package services

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// API constants for the resume parsing service

const apiURL = "https://api.apilayer.com/resume_parser/upload"

// UploadResume handles the uploading of resumes to the third-party API, processes the response, and stores the data in MongoDB.
func UploadResume(filePath string, userID string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a multipart form to send the file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return fmt.Errorf("failed to create form file: %v", err)
	}
	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	writer.Close()

	// Create the POST request to the third-party API
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY") // Get the MongoDB URI from the .env file
	req.Header.Set("apikey", apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {

		if resp.StatusCode == 429 {
			retryAfter := resp.Header.Get("Retry-After")
			log.Println("Too many requests. Retry after:", retryAfter)
			// Convert `retryAfter` to time.Duration and wait
			return fmt.Errorf("received non-200 response: %s\n Retry after:%s", resp.Status, retryAfter)
		}
	}

	// Read the API response
	resumeData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Log the raw API response for debugging purposes
	log.Println("API Response: ", string(resumeData))

	// Parse the API response
	var profileAPIResponse models.ProfileAPIResponse
	if err = json.Unmarshal(resumeData, &profileAPIResponse); err != nil {
		return fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	// Log parsed data
	log.Println("Parsed Name: ", profileAPIResponse.Name)
	log.Println("Parsed Email: ", profileAPIResponse.Email)
	log.Println("Parsed Phone: ", profileAPIResponse.Phone)

	// Map the parsed response to the Profile model for MongoDB storage
	profile := models.Profile{
		ResumeFileAddress: filePath,                                         // Store the resume file path
		Skills:            extractSkills(profileAPIResponse.Skills),         // Extract skills as a string
		Education:         extractEducation(profileAPIResponse.Education),   // Extract education as a string
		Experience:        extractExperience(profileAPIResponse.Experience), // Extract experience as a string
		Phone:             profileAPIResponse.Phone,
		Email:             profileAPIResponse.Email,
		Name:              profileAPIResponse.Name,
	}

	// Get the MongoDB collection and update the profile data
	collection := utils.GetCollection("profiles")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},   // Update the existing profile with this user's ID
		bson.M{"$set": profile}, // Set the new profile data
	)
	if err != nil {
		return fmt.Errorf("failed to insert data into MongoDB: %v", err)
	}

	log.Println("Resume data successfully inserted into MongoDB for user:", userID)

	return nil
}

// Helper function to extract skills as a string
func extractSkills(skills []string) string {
	if len(skills) == 0 {
		return ""
	}
	return fmt.Sprintf("%v", skills)
}

// Helper function to extract education as a string
func extractEducation(education []models.EducationAPIResponse) string {
	educationList := ""
	for _, edu := range education {
		educationList += edu.Name + " "
	}
	return educationList
}

// Helper function to extract experience as a string
func extractExperience(experience []models.ExperienceAPIResponse) string {
	experienceList := ""
	for _, exp := range experience {
		experienceList += exp.Name + " "
	}
	return experienceList
}
