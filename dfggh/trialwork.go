package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main(){
	s := server.NewMCPServer(
		"math-server",
		"1.0.0",
		server.WithLogging(),
	)

	s.AddTool(mcp.NewTool("add",
	))

	
}