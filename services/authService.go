package services

import (
	"RecruitmentManagementSystem/models"
	"RecruitmentManagementSystem/utils"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	UserType        string `json:"user_type"` // Admin or Applicant
	Password        string `json:"password"`
	ProfileHeadline string `json:"profile_headline"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(userRequest UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:            userRequest.Name,
		Email:           userRequest.Email,
		Address:         userRequest.Address,
		UserType:        userRequest.UserType,
		PasswordHash:    string(hashedPassword),
		ProfileHeadline: userRequest.ProfileHeadline,
	}

	collection := utils.Client.Database("recruitment_management").Collection("users")
	_, err = collection.InsertOne(context.TODO(), user)
	return err
}

func Login(userRequest LoginRequest) (string, error) {
	collection := utils.Client.Database("recruitment_management").Collection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": userRequest.Email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userRequest.Password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	return tokenString, err
}
