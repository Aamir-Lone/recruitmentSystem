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

//Profile structure
type Profile struct {
	ResumeFileAddress string       `bson:"resumeFileAddress"`
	Skills            []string     `bson:"skills"`
	Education         []Education  `bson:"education"`  // Updated to slice of Education structs
	Experience        []Experience `bson:"experience"` // Updated to slice of Experience structs
	Phone             string       `bson:"phone"`
}

type Education struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"` // This is an array of strings
}

type Experience struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
}
