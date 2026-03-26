package middleware

import (
	"aioz-pin-mcp-server/pkg/cache"
	"aioz-pin-mcp-server/util"
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	gocache "github.com/patrickmn/go-cache"
)

func AuthMiddleware() func(server.ToolHandlerFunc) server.ToolHandlerFunc {
	return func(next server.ToolHandlerFunc) server.ToolHandlerFunc {
		return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			authcache := cache.GetAuthCache()

			args, ok := req.Params.Arguments.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("Invalid arguments")
			}

			jwtToken, _ := args["jwtToken"].(string)

			if val, found := authcache.Get(jwtToken); found {
				req.Params.Arguments.(map[string]interface{})["client"] = val
				return next(ctx, req)
			}

			resp, err := util.MakeRequest("users/me", "GET", "", jwtToken)
			if err != nil {
				return nil, fmt.Errorf("Authentication failed: %v", err)
			}

			user, ok := resp["user"].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("Invalid response format: missing data")
			}

			isVerified, ok := user["is_verified"].(bool)
			if !ok || !isVerified {
				return nil, fmt.Errorf("User verification status is empty or invalid")
			}

			authcache.Set(jwtToken, jwtToken, gocache.DefaultExpiration)
			req.Params.Arguments.(map[string]interface{})["jwtToken"] = jwtToken

			return next(ctx, req)
		}
	}
}
