package model

import "fmt"

type HistoryUsageInput struct {
	Offset int
	Limit  int
	AuthInput
}

func ValidateHistoryUsageInput(args map[string]interface{}) (*HistoryUsageInput, error) {
	input := &HistoryUsageInput{}

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

	if input.Limit <= 0 {
		return &HistoryUsageInput{}, fmt.Errorf("limit must be greater than 0")
	}
	if input.Offset < 0 {
		return &HistoryUsageInput{}, fmt.Errorf("offset must be greater than or equal to 0")
	}

	return input, nil
}
