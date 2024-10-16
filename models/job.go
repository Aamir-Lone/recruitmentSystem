package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Title             string             `bson:"title"`
	Description       string             `bson:"description"`
	PostedOn          string             `bson:"posted_on"`
	TotalApplications int                `bson:"total_applications"`
	CompanyName       string             `bson:"company_name"`
	PostedBy          User               `bson:"posted_by"`
}
