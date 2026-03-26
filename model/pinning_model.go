package model

import "fmt"

type PinByCIDInput struct {
	HashToPin    string
	MetadataName string
	APIKey       string
	SecretKey    string
}

type GetPinDetailsInput struct {
	PinID     string
	APIKey    string
	SecretKey string
}

type ListPinsInput struct {
	Offset    int
	Limit     int
	Pinned    bool
	SortBy    string
	SortOrder string
	APIKey    string
	SecretKey string
}

type UnpinInput struct {
	PinID     string
	APIKey    string
	SecretKey string
}

func ValidatePinByCIDInput(args map[string]interface{}) (*PinByCIDInput, error) {
	input := &PinByCIDInput{}

	if hashToPin, ok := args["hashToPin"].(string); ok {
		input.HashToPin = hashToPin
	}
	if metadataName, ok := args["metadataName"].(string); ok {
		input.MetadataName = metadataName
	}
	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}
	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.HashToPin == "" {
		return &PinByCIDInput{}, fmt.Errorf("hashToPin is required")
	}
	if input.APIKey == "" {
		return &PinByCIDInput{}, fmt.Errorf("pinningApiKey is required")
	}
	if input.SecretKey == "" {
		return &PinByCIDInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}

func ValidateGetPinDetailsInput(args map[string]interface{}) (*GetPinDetailsInput, error) {
	input := &GetPinDetailsInput{}

	if pinID, ok := args["pinId"].(string); ok {
		input.PinID = pinID
	}
	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}
	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.PinID == "" {
		return &GetPinDetailsInput{}, fmt.Errorf("pinId is required")
	}
	if input.APIKey == "" {
		return &GetPinDetailsInput{}, fmt.Errorf("pinningApiKey is required")
	}
	if input.SecretKey == "" {
		return &GetPinDetailsInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}

func ValidateListPinsInput(args map[string]interface{}) (*ListPinsInput, error) {
	input := &ListPinsInput{
		Offset:    0,
		Limit:     10,
		Pinned:    true,
		SortBy:    "name",
		SortOrder: "ASC",
	}

	if offset, ok := args["offset"].(float64); ok {
		input.Offset = int(offset)
	}
	if limit, ok := args["limit"].(float64); ok {
		input.Limit = int(limit)
	}
	if pinned, ok := args["pinned"].(bool); ok {
		input.Pinned = pinned
	}
	if sortBy, ok := args["sortBy"].(string); ok && sortBy != "" {
		input.SortBy = sortBy
	}
	if sortOrder, ok := args["sortOrder"].(string); ok && sortOrder != "" {
		input.SortOrder = sortOrder
	}
	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}
	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.APIKey == "" {
		return &ListPinsInput{}, fmt.Errorf("pinningApiKey is required")
	}
	if input.SecretKey == "" {
		return &ListPinsInput{}, fmt.Errorf("pinningSecretKey is required")
	}
	if input.Limit <= 0 {
		return &ListPinsInput{}, fmt.Errorf("limit must be greater than 0")
	}
	if input.Offset < 0 {
		return &ListPinsInput{}, fmt.Errorf("offset must be greater than or equal to 0")
	}

	return input, nil
}

func ValidateUnpinInput(args map[string]interface{}) (*UnpinInput, error) {
	input := &UnpinInput{}

	if pinID, ok := args["pinId"].(string); ok {
		input.PinID = pinID
	}
	if apiKey, ok := args["pinningApiKey"].(string); ok {
		input.APIKey = apiKey
	}
	if secretKey, ok := args["pinningSecretKey"].(string); ok {
		input.SecretKey = secretKey
	}

	if input.PinID == "" {
		return &UnpinInput{}, fmt.Errorf("pinId is required")
	}
	if input.APIKey == "" {
		return &UnpinInput{}, fmt.Errorf("pinningApiKey is required")
	}
	if input.SecretKey == "" {
		return &UnpinInput{}, fmt.Errorf("pinningSecretKey is required")
	}

	return input, nil
}
