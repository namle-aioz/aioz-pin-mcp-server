package handler

import (
	"aioz-pin-mcp-server/middleware"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterHandlers(mcpServer *server.MCPServer) {

	requiredAuth := true

	generateAPIKeyTool := newAIOZTool(
		"generate-api-key",
		`Generate an AIOZ Pin API key. The key must be identified as an admin key or a regular key. 
		Admin keys have full access to all scopes, while non-admin keys require at least one specific permission to be enabled.`,
		requiredAuth,

		mcp.WithString(
			"keyName",
			mcp.Description("Name of the API key"),
			mcp.Required(),
		),

		mcp.WithBoolean("admin", mcp.Description("If true, grants full access to all scopes"), mcp.Required()),
		mcp.WithBoolean("pinList", mcp.Description("Allow listing pins"), mcp.Required()),
		mcp.WithBoolean("nftList", mcp.Description("Allow listing NFTs"), mcp.Required()),
		mcp.WithBoolean("unpin", mcp.Description("Allow unpinning content"), mcp.Required()),
		mcp.WithBoolean("pinByHash", mcp.Description("Allow pinning by hash"), mcp.Required()),
		mcp.WithBoolean("pinFileToIPFS", mcp.Description("Allow uploading files to IPFS"), mcp.Required()),
		mcp.WithBoolean("unpinNFT", mcp.Description("Allow unpinning NFTs"), mcp.Required()),
		mcp.WithBoolean("pinNFTToIPFS", mcp.Description("Allow uploading NFTs to IPFS"), mcp.Required()),
	)

	getListAPIKeys := newAIOZTool(
		"get-list-api-keys",
		`Retrieve a list of all AIOZ Pin API keys.`,
		requiredAuth,
	)

	deleteAPIKey := newAIOZTool(
		"delete-api-key",
		`Delete an AIOZ Pin API key.`,
		requiredAuth,
		mcp.WithString(
			"keyId",
			mcp.Description("ID of the API key to delete"),
			mcp.Required(),
		),
	)

	pinFilesOrDirectoryTool := newAIOZTool(
		"pin-files-or-directory",
		`This API allows you to pin a file to IPFS using the provided pinning API key and secret key. You can specify the file to be pinned by providing a publicly accessible URL. The server will download the file from the URL and then upload it to IPFS using the pinning service.`,
		false,
		mcp.WithString(
			"fileUrl",
			mcp.Description("Public downloadable file URL"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	pinByCIDTool := newAIOZTool(
		"pin-by-cid",
		`Pin content by CID hash using pinning API key and secret key.`,
		false,
		mcp.WithString(
			"hashToPin",
			mcp.Description("CID hash to pin"),
			mcp.Required(),
		),
		mcp.WithString(
			"metadataName",
			mcp.Description("Optional name metadata for pinned content"),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	getPinDetailsTool := newAIOZTool(
		"get-pin-details",
		`Get pin details by pin ID.`,
		false,
		mcp.WithString(
			"pinId",
			mcp.Description("Pin ID to fetch details"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	listPinsTool := newAIOZTool(
		"list-pins",
		`List pins with pagination and sorting options.`,
		false,
		mcp.WithNumber(
			"offset",
			mcp.Description("Pagination offset, default 0"),
		),
		mcp.WithNumber(
			"limit",
			mcp.Description("Pagination limit, default 10"),
		),
		mcp.WithBoolean(
			"pinned",
			mcp.Description("Filter by pinned state, default true"),
		),
		mcp.WithString(
			"sortBy",
			mcp.Description("Sort field, default name"),
		),
		mcp.WithString(
			"sortOrder",
			mcp.Description("Sort order ASC or DESC, default ASC"),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	unpinTool := newAIOZTool(
		"unpin-file",
		`Remove pinned file by pin ID.`,
		false,
		mcp.WithString(
			"pinId",
			mcp.Description("Pin ID to remove"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	getHistoryUsageDataTool := newAIOZTool(
		"get-history-usage-data",
		`Get history usage data details.`,
		false,
		mcp.WithString(
			"offset",
			mcp.Description("Pagination offset default 0"),
			mcp.Required(),
		),
		mcp.WithString(
			"limit",
			mcp.Description("Pagination limit default 10"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	getTopUpTool := newAIOZTool(
		"get-top-up",
		`Get top-up data details.`,
		false,
		mcp.WithString(
			"offset",
			mcp.Description("Pagination offset default 0"),
			mcp.Required(),
		),
		mcp.WithString(
			"limit",
			mcp.Description("Pagination limit default 10"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	getMonthUsageDataTool := newAIOZTool(
		"get-month-usage-data",
		`Get month usage data details.`,
		false,
		mcp.WithString(
			"offset",
			mcp.Description("Pagination offset default 0"),
			mcp.Required(),
		),
		mcp.WithString(
			"limit",
			mcp.Description("Pagination limit default 10"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningApiKey",
			mcp.Description("AIOZ Pinning API key"),
			mcp.Required(),
		),
		mcp.WithString(
			"pinningSecretKey",
			mcp.Description("AIOZ Pinning secret key"),
			mcp.Required(),
		),
	)

	auth := middleware.AuthMiddleware()

	mcpServer.AddTool(generateAPIKeyTool, auth(HandleGenerateAPIKey))
	mcpServer.AddTool(getListAPIKeys, auth(HandleGetListAPIKeys))
	mcpServer.AddTool(deleteAPIKey, auth(HandleDeleteAPIKey))
	mcpServer.AddTool(pinFilesOrDirectoryTool, HandlePinFilesOrDirectory)
	mcpServer.AddTool(pinByCIDTool, HandlePinByCID)
	mcpServer.AddTool(getPinDetailsTool, HandleGetPinDetails)
	mcpServer.AddTool(listPinsTool, HandleListPins)
	mcpServer.AddTool(unpinTool, HandleUnpin)
	mcpServer.AddTool(getHistoryUsageDataTool, HandleGetHistoryUsageData)
	mcpServer.AddTool(getTopUpTool, HandleGetTopUp)
	mcpServer.AddTool(getMonthUsageDataTool, HandleGetMonthUsageData)
}

func newAIOZTool(name string, description string, requiredAuth bool, params ...mcp.ToolOption) mcp.Tool {
	baseParams := []mcp.ToolOption{
		mcp.WithDescription(description),
	}

	if requiredAuth {
		baseParams = append(baseParams,
			mcp.WithString(
				"jwtToken",
				mcp.Description("JWT token for authorization"),
				mcp.Required(),
			),
		)
	}

	baseParams = append(baseParams, params...)
	return mcp.NewTool(name, baseParams...)
}
