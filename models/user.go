package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Email           string             `bson:"email"`
	Address         string             `bson:"address"`
	UserType        string             `bson:"userType"` // Applicant/Admin
	PasswordHash    string             `bson:"passwordHash"`
	ProfileHeadline string             `bson:"profileHeadline"`
	Profile         Profile            `bson:"profile"`
}

// Profile structure
type Profile struct {
	ResumeFileAddress string `bson:"resumeFileAddress"`
	Skills            string `bson:"skills"`
	Education         string `bson:"education"`
	Experience        string `bson:"experience"`
	Phone             string `bson:"phone"`
}
