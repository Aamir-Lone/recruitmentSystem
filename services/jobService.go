package services

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateJob(job models.Job) (*models.Job, error) {
	job.ID = primitive.NewObjectID()
	job.PostedOn = primitive.NewDateTimeFromTime(time.Now())
	collection := utils.GetCollection("jobs")
	_, err := collection.InsertOne(context.TODO(), job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
