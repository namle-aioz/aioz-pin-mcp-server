package handler

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPinningTools(mcpServer *server.MCPServer) {
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

	mcpServer.AddTool(pinFilesOrDirectoryTool, HandlePinFilesOrDirectory)
	mcpServer.AddTool(pinByCIDTool, HandlePinByCID)
	mcpServer.AddTool(getPinDetailsTool, HandleGetPinDetails)
	mcpServer.AddTool(listPinsTool, HandleListPins)
	mcpServer.AddTool(unpinTool, HandleUnpin)
}
