package handler

import (
	"aioz-pin-mcp-server/middleware"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterHandlers(mcpServer *server.MCPServer) {
	auth := middleware.AuthMiddleware()

	registerAPIKeyTools(mcpServer, auth)
	registerPinningTools(mcpServer)
	registerUsageTools(mcpServer)
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
