package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Title             string             `bson:"title"`
	Description       string             `bson:"description"`
	PostedOn          primitive.DateTime `bson:"postedOn"`
	TotalApplications int                `bson:"totalApplications"`
	CompanyName       string             `bson:"companyName"`
	PostedBy          primitive.ObjectID `bson:"postedBy"`
}
