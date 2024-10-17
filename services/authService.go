package services

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	collection := utils.GetCollection("users")
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("incorrect password")
	}
	return &user, nil
}

func CreateUser(user models.User) (*models.User, error) {
	user.PasswordHash, _ = HashPassword(user.PasswordHash)
	collection := utils.GetCollection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
