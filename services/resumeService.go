package services

import (
	"RecruitmentManagementSystem/models"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const apiKey = "0bWeisRWoLj3UdXt3MXMSMWptYFIpQfS"
const apiURL = "https://api.apilayer.com/resume_parser/upload"

func UploadResume(filePath string) (*models.ProfileAPIResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", file.Name())
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resumeData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profile models.ProfileAPIResponse
	err = json.Unmarshal(resumeData, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
