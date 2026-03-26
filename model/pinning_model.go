package model

import "fmt"

type PinByCIDInput struct {
	HashToPin    string
	MetadataName string
	AuthInput
}

type GetPinDetailsInput struct {
	PinID string
	AuthInput
}

type ListPinsInput struct {
	Offset    int
	Limit     int
	Pinned    bool
	SortBy    string
	SortOrder string
	AuthInput
}

type UnpinInput struct {
	PinID string
	AuthInput
}

func ValidatePinByCIDInput(args map[string]interface{}) (*PinByCIDInput, error) {
	input := &PinByCIDInput{}

	authInput, err := ValidateAuthInput(args)
	if err != nil {
		return nil, err
	}

	input.AuthInput = *authInput

	if hashToPin, ok := args["hashToPin"].(string); ok {
		input.HashToPin = hashToPin
	}
	if metadataName, ok := args["metadataName"].(string); ok {
		input.MetadataName = metadataName
	}
	if input.HashToPin == "" {
		return &PinByCIDInput{}, fmt.Errorf("hashToPin is required")
	}

	return input, nil
}

func ValidateGetPinDetailsInput(args map[string]interface{}) (*GetPinDetailsInput, error) {
	input := &GetPinDetailsInput{}

	authInput, err := ValidateAuthInput(args)
	if err != nil {
		return nil, err
	}

	input.AuthInput = *authInput

	if pinID, ok := args["pinId"].(string); ok {
		input.PinID = pinID
	}
	if input.PinID == "" {
		return &GetPinDetailsInput{}, fmt.Errorf("pinId is required")
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

	authInput, err := ValidateAuthInput(args)
	if err != nil {
		return nil, err
	}

	input.AuthInput = *authInput

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

	authInput, err := ValidateAuthInput(args)
	if err != nil {
		return nil, err
	}

	input.AuthInput = *authInput

	if pinID, ok := args["pinId"].(string); ok {
		input.PinID = pinID
	}

	if input.PinID == "" {
		return &UnpinInput{}, fmt.Errorf("pinId is required")
	}

	return input, nil
}
