package model

import (
	"fmt"
)

type GenerateAPIKeyInput struct {
	KeyName       string
	Admin         bool
	PinList       bool
	NFTList       bool
	Unpin         bool
	PinByHash     bool
	PinFileToIPFS bool
	UnpinNFT      bool
	PinNFTToIPFS  bool
}

func ValidateGenerateAPIKeyInput(args map[string]interface{}) (*GenerateAPIKeyInput, error) {
	input := &GenerateAPIKeyInput{}

	if keyName, ok := args["keyName"].(string); ok {
		input.KeyName = keyName
	}

	if isAdmin, ok := args["admin"].(bool); ok {
		input.Admin = isAdmin
	}

	if pinList, ok := args["pinList"].(bool); ok {
		input.PinList = pinList
	}
	if nftList, ok := args["nftList"].(bool); ok {
		input.NFTList = nftList
	}
	if unpin, ok := args["unpin"].(bool); ok {
		input.Unpin = unpin
	}
	if pinByHash, ok := args["pinByHash"].(bool); ok {
		input.PinByHash = pinByHash
	}
	if pinFileToIPFS, ok := args["pinFileToIPFS"].(bool); ok {
		input.PinFileToIPFS = pinFileToIPFS
	}
	if unpinNFT, ok := args["unpinNFT"].(bool); ok {
		input.UnpinNFT = unpinNFT
	}
	if pinNFTToIPFS, ok := args["pinNFTToIPFS"].(bool); ok {
		input.PinNFTToIPFS = pinNFTToIPFS
	}

	if input.KeyName == "" {
		return &GenerateAPIKeyInput{}, fmt.Errorf("keyName is required")
	}

	if input.Admin {
		return input, nil
	}

	if !input.PinList &&
		!input.NFTList &&
		!input.Unpin &&
		!input.PinByHash &&
		!input.PinFileToIPFS &&
		!input.UnpinNFT &&
		!input.PinNFTToIPFS {

		return &GenerateAPIKeyInput{}, fmt.Errorf("at least one permission must be enabled when admin is false")
	}

	return input, nil
}
