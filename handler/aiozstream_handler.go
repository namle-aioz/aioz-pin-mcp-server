package handler

import (
	"aioz-pin-mcp-server/model"
	"aioz-pin-mcp-server/tool"
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleGenerateAPIKey(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	keyName, ok := args["keyName"].(string)
	if !ok {
		return mcp.NewToolResultError("Key name parameter required"), nil
	}

	input, err := model.ValidateGenerateAPIKeyInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	client, ok := args["jwtToken"].(string)
	if !ok {
		return mcp.NewToolResultError("Unauthorized"), nil
	}

	key, err := tool.GenerateAPIKey(ctx, client, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf(
		"Key created successfully to AIOZ Stream with name: %+v and key: %+v",
		keyName, key,
	)

	return mcp.NewToolResultText(result), nil
}

func HandleGetListAPIKeys(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	client, ok := args["jwtToken"].(string)
	if !ok {
		return mcp.NewToolResultError("Unauthorized"), nil
	}

	keys, err := tool.GetListAPIKeys(ctx, client)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf(
		"API Keys retrieved successfully: %+v",
		keys,
	)

	return mcp.NewToolResultText(result), nil
}

func HandleDeleteAPIKey(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	client, ok := args["jwtToken"].(string)
	if !ok {
		return mcp.NewToolResultError("Unauthorized"), nil
	}

	keyId, ok := args["keyId"].(string)
	if !ok {
		return mcp.NewToolResultError("Key ID parameter required"), nil
	}

	_, err := tool.DeleteAPIKeyById(ctx, client, keyId)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf(
		"API Key deleted successfully: %+v",
		keyId,
	)

	return mcp.NewToolResultText(result), nil
}

func HandlePinFilesOrDirectory(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidatePinFilesOrDirectoryInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.PinFilesOrDirectory(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Pinned successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandlePinByCID(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidatePinByCIDInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.PinByCID(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Pinned by CID successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleGetPinDetails(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateGetPinDetailsInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.GetPinDetails(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Pin details retrieved successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleListPins(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateListPinsInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.ListPins(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Pins listed successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleUnpin(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateUnpinInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.Unpin(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Unpinned successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleGetHistoryUsageData(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateHistoryUsageInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.GetHistoryUsageData(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("History usage data retrieved successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleGetTopUp(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateHistoryUsageInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.GetTopUp(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Top-up data retrieved successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}

func HandleGetMonthUsageData(
	ctx context.Context,
	req mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	args, ok := req.Params.Arguments.(map[string]interface{})
	if !ok {
		return mcp.NewToolResultError("Invalid arguments"), nil
	}

	input, err := model.ValidateHistoryUsageInput(args)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := tool.GetMonthUsageData(ctx, input)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result := fmt.Sprintf("Month usage data retrieved successfully: %+v", resp)

	return mcp.NewToolResultText(result), nil
}
