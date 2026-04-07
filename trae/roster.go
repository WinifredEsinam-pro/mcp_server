package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var (
	soldiers []string
	police   []string
	mu       sync.Mutex
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func removeItems(slice []string, items []string) []string {
	toRemove := make(map[string]struct{}, len(items))
	for _, n := range items {
		toRemove[n] = struct{}{}
	}
	out := make([]string, 0, len(slice))
	for _, v := range slice {
		if _, ok := toRemove[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

func RegisterRosterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("update_roster",
		mcp.WithDescription("Add, update, or remove names in soldiers or police lists"),
		mcp.WithString("group",
			mcp.Required(),
			mcp.Description("Which group to update (soldiers or police)"),
			mcp.Enum("soldiers", "police"),
		),
		mcp.WithString("mode",
			mcp.Required(),
			mcp.Description("The operation: append, remove, replace, or clear"),
			mcp.Enum("append", "remove", "replace", "clear"),
		),
		mcp.WithArray("names",
			mcp.Description("List of names for append, remove, or replace"),
			mcp.WithStringItems(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		group, err := request.RequireString("group")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid 'group': %v", err)), nil
		}
		mode, err := request.RequireString("mode")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid 'mode': %v", err)), nil
		}
		names := request.GetStringSlice("names", []string{})

		mu.Lock()
		defer mu.Unlock()

		var targetList *[]string
		switch group {
		case "soldiers":
			targetList = &soldiers
		case "police":
			targetList = &police
		default:
			return mcp.NewToolResultError("invalid group"), nil
		}

		switch mode {
		case "append":
			for _, n := range names {
				if !contains(*targetList, n) {
					*targetList = append(*targetList, n)
				}
			}
		case "remove":
			if len(names) == 0 {
				return mcp.NewToolResultError("names required for remove"), nil
			}
			*targetList = removeItems(*targetList, names)
		case "replace":
			*targetList = append([]string(nil), names...)
		case "clear":
			*targetList = nil
		default:
			return mcp.NewToolResultError("invalid mode"), nil
		}

		return mcp.NewToolResultJSON(map[string]any{
			"group": group,
			"mode":  mode,
			"list":  *targetList,
		})
	})
}
