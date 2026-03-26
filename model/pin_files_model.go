package model

import (
	"fmt"
	"net/url"
)

type PinFilesOrDirectoryInput struct {
	FileURL   string
	APIKey    string
	SecretKey string
}

func ValidatePinFilesOrDirectoryInput(args map[string]interface{}) (*PinFilesOrDirectoryInput, error) {
	input := &PinFilesOrDirectoryInput{}

	if fileURL, ok := args["fileUrl"].(string); ok {
		input.FileURL = fileURL
	}

	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}

	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.FileURL == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("fileUrl is required")
	}

	parsedURL, err := url.Parse(input.FileURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("fileUrl is invalid")
	}

	if input.APIKey == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("pinningApiKey is required")
	}

	if input.SecretKey == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}
