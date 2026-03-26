package tool

import (
	"aioz-pin-mcp-server/model"
	"aioz-pin-mcp-server/util"
	"context"
	"fmt"
	"net/url"
	"os"
)

func GenerateAPIKey(ctx context.Context, jwtToken string, input *model.GenerateAPIKeyInput) (map[string]interface{}, error) {
	var payload string
	if input.Admin {
		payload = fmt.Sprintf(`{"name":%q,"scopes":{"admin":true}}`, input.KeyName)
	} else {
		payload = fmt.Sprintf(`{"name":%q,"scopes":{"admin":%t,"data":{"pin_list":%t,"nft_list":%t},"pinning":{"unpin":%t,"pin_by_hash":%t,"pin_file_to_ipfs":%t},"pin_nft":{"unpin_nft":%t,"pin_nft_to_ipfs":%t}}}`,
			input.KeyName,
			input.Admin,
			input.PinList,
			input.NFTList,
			input.Unpin,
			input.PinByHash,
			input.PinFileToIPFS,
			input.UnpinNFT,
			input.PinNFTToIPFS,
		)
	}

	resp, err := util.MakeRequest("apiKeys", "POST", payload, jwtToken)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetListAPIKeys(ctx context.Context, jwtToken string) (map[string]interface{}, error) {
	resp, err := util.MakeRequest("apiKeys/list", "GET", "", jwtToken)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteAPIKeyById(ctx context.Context, jwtToken string, keyId string) (map[string]interface{}, error) {
	resp, err := util.MakeRequest(fmt.Sprintf("apiKeys/%s", keyId), "DELETE", "", jwtToken)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PinFilesOrDirectory(ctx context.Context, input *model.PinFilesOrDirectoryInput) (map[string]interface{}, error) {
	tempPath, fileName, contentType, err := util.GetFileFromURL(ctx, input.FileURL)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempPath)

	uploadResp, err := util.UploadPinningFile("pinning/", tempPath, input.APIKey, input.SecretKey)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"fileUrl":     input.FileURL,
		"fileName":    fileName,
		"contentType": contentType,
		"response":    uploadResp,
	}, nil
}

func PinByCID(ctx context.Context, input *model.PinByCIDInput) (map[string]interface{}, error) {
	var payload string
	if input.MetadataName != "" {
		payload = fmt.Sprintf(`{"hash_to_pin":%q,"metadata":{"name":%q}}`, input.HashToPin, input.MetadataName)
	} else {
		payload = fmt.Sprintf(`{"hash_to_pin":%q}`, input.HashToPin)
	}

	return util.MakePinningRequest("pinning/pinByHash", "POST", payload, input.APIKey, input.SecretKey)
}

func GetPinDetails(ctx context.Context, input *model.GetPinDetailsInput) (map[string]interface{}, error) {
	return util.MakePinningRequest(fmt.Sprintf("pinning/%s", input.PinID), "GET", "", input.APIKey, input.SecretKey)
}

func ListPins(ctx context.Context, input *model.ListPinsInput) (map[string]interface{}, error) {
	query := fmt.Sprintf(
		"offset=%d&limit=%d&pinned=%t&sortBy=%s&sortOrder=%s",
		input.Offset,
		input.Limit,
		input.Pinned,
		url.QueryEscape(input.SortBy),
		url.QueryEscape(input.SortOrder),
	)

	return util.MakePinningRequest(fmt.Sprintf("pinning/pins/?%s", query), "GET", "", input.APIKey, input.SecretKey)
}

func Unpin(ctx context.Context, input *model.UnpinInput) (map[string]interface{}, error) {
	return util.MakePinningRequest(fmt.Sprintf("pinning/unpin/%s", input.PinID), "DELETE", "", input.APIKey, input.SecretKey)
}

func GetHistoryUsageData(ctx context.Context, input *model.HistoryUsageInput) (map[string]interface{}, error) {
	return util.MakePinningRequest(fmt.Sprintf("billing/historyUsage?offset=%d&limit=%d", input.Offset, input.Limit), "GET", "", input.APIKey, input.SecretKey)
}

func GetTopUp(ctx context.Context, input *model.HistoryUsageInput) (map[string]interface{}, error) {
	return util.MakePinningRequest(fmt.Sprintf("billing/topUp?offset=%d&limit=%d", input.Offset, input.Limit), "GET", "", input.APIKey, input.SecretKey)
}

func GetMonthUsageData(ctx context.Context, input *model.HistoryUsageInput) (map[string]interface{}, error) {
	return util.MakePinningRequest(fmt.Sprintf("billing/thisMonthUsage?offset=%d&limit=%d", input.Offset, input.Limit), "GET", "", input.APIKey, input.SecretKey)
}
