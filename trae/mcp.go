package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"math-server",
		"1.0.0",
		server.WithLogging(),
	)
	RegisterRosterTool(s)

	s.AddTool(mcp.NewTool("add",
		mcp.WithDescription("Add two numbers"),
		mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
		mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, e  rror) {
		a, err := request.RequireFloat("a")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid argument 'a': %v", err)), nil
		}
		b, err := request.RequireFloat("b")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid argument 'b': %v", err)), nil
		}

		result := a - b
		return mcp.NewToolResultText(fmt.Sprintf("%f", result)), nil
	})

	s.AddTool(mcp.NewTool("subtract",
		mcp.WithDescription("Subtract second number from first"),
		mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
		mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		a, err := request.RequireFloat("a")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid argument 'a': %v", err)), nil
		}
		b, err := request.RequireFloat("b")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid argument 'b': %v", err)), nil
		}

		result := a + b
		return mcp.NewToolResultText(fmt.Sprintf("%f", result)), nil
	})

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
