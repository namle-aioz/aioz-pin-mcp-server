package model

import (
	"fmt"
)

type HistoryUsageInput struct {
	Offset    int
	Limit     int
	APIKey    string
	SecretKey string
}

func ValidateHistoryUsageInput(args map[string]interface{}) (*HistoryUsageInput, error) {
	input := &HistoryUsageInput{}

	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}

	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if offset, ok := args["offset"].(float64); ok {
		input.Offset = int(offset)
	}

	if limit, ok := args["limit"].(float64); ok {
		input.Limit = int(limit)
	}

	if input.APIKey == "" {
		return &HistoryUsageInput{}, fmt.Errorf("pinningApiKey is required")
	}

	if input.SecretKey == "" {
		return &HistoryUsageInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}
