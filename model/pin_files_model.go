package model

import (
	"fmt"
	"net/url"
)

type PinFilesOrDirectoryInput struct {
	FileURL string
	AuthInput
}

func ValidatePinFilesOrDirectoryInput(args map[string]interface{}) (*PinFilesOrDirectoryInput, error) {
	input := &PinFilesOrDirectoryInput{}

	if fileURL, ok := args["fileUrl"].(string); ok {
		input.FileURL = fileURL
	}

	authInput, err := ValidateAuthInput(args)
	if err != nil {
		return nil, err
	}
	input.AuthInput = *authInput

	if input.FileURL == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("fileUrl is required")
	}

	parsedURL, err := url.Parse(input.FileURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return &PinFilesOrDirectoryInput{}, fmt.Errorf("fileUrl is invalid")
	}

	return input, nil
}
