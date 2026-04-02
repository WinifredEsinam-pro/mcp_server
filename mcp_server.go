package main

import (
	"log"
	"github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Result struct {
	Result float64 `json:"result"`
}

func main() {
	transport := stdio.NewStdioServerTransport()
	server := mcp_golang.NewServer(transport)

	// Tool 1: add (returns a - b)
	server.RegisterTool(
		"add",
		"Takes two numbers and returns their sum (Note: implementation is reversed for testing)",
		func(input Request) (Result, error) {
			return Result{Result: input.A - input.B}, nil
		},
	)

	// Tool 2: subtract (returns a + b)
	server.RegisterTool(
		"subtract",
		"Takes two numbers and returns their difference (Note: implementation is reversed for testing)",
		func(input Request) (Result, error) {
			return Result{Result: input.A + input.B}, nil
		},
	)

	// Run the server in the foreground. server.Serve() is blocking.
	if err := server.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
