package handler

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerAPIKeyTools(mcpServer *server.MCPServer, auth func(server.ToolHandlerFunc) server.ToolHandlerFunc) {
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

	mcpServer.AddTool(generateAPIKeyTool, auth(HandleGenerateAPIKey))
	mcpServer.AddTool(getListAPIKeys, auth(HandleGetListAPIKeys))
	mcpServer.AddTool(deleteAPIKey, auth(HandleDeleteAPIKey))
}
