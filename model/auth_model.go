package model

import "fmt"

type AuthInput struct {
	APIKey    string
	SecretKey string
}

func ValidateAuthInput(args map[string]interface{}) (*AuthInput, error) {
	input := &AuthInput{}

	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}

	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.APIKey == "" {
		return &AuthInput{}, fmt.Errorf("pinningApiKey is required")
	}

	if input.SecretKey == "" {
		return &AuthInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}
