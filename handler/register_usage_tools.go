package handler

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerUsageTools(mcpServer *server.MCPServer) {
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

	mcpServer.AddTool(getHistoryUsageDataTool, HandleGetHistoryUsageData)
	mcpServer.AddTool(getTopUpTool, HandleGetTopUp)
	mcpServer.AddTool(getMonthUsageDataTool, HandleGetMonthUsageData)
}
