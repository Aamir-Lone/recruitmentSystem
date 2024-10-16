package services

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type JobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyName string `json:"company_name"`
}

func CreateJob(jobRequest JobRequest) error {
	job := models.Job{
		Title:             jobRequest.Title,
		Description:       jobRequest.Description,
		PostedOn:          time.Now().Format("2006-01-02"),
		TotalApplications: 0,
		CompanyName:       jobRequest.CompanyName,
		PostedBy:          models.User{}, // You might want to associate this with the logged-in user
	}

	collection := utils.Client.Database("recruitment_management").Collection("jobs")
	_, err := collection.InsertOne(context.TODO(), job)
	return err
}

func GetJobByID(jobID string) (models.Job, error) {
	var job models.Job
	collection := utils.Client.Database("recruitment_management").Collection("jobs")
	err := collection.FindOne(context.TODO(), bson.M{"_id": jobID}).Decode(&job)
	return job, err
}

// GetAllApplicants retrieves all applicants from the database
func GetAllApplicants() ([]models.User, error) {
	collection := utils.Client.Database("your_database_name").Collection("users")

	var applicants []models.User
	filter := bson.M{"user_type": "Applicant"} // Adjust as necessary
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var applicant models.User
		if err := cursor.Decode(&applicant); err != nil {
			return nil, err
		}
		applicants = append(applicants, applicant)
	}

	return applicants, nil
}
