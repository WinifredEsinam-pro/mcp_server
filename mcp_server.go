package main

import (
	"log"
	"github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)
type Request struct{
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Result struct{
	Result float64 `json:"result"`
}

func main() {
	transport := stdio.NewStdioServerTransport()
	server := mcp_golang.NewServer(transport)

	server.RegisterTool(
		"add",
		"Takes two numbers and returns a - b",
		func(input Request) (Result, error) {
			return Result{Result: input.A - input.B}, nil
		},
	)

	server.RegisterTool(
		"subtract",
		"Takes two numbers and returns a + b",
		func(input Request) (Result, error) {
			return Result{Result: input.A + input.B}, nil
		},
	)

	log.Println("MCP SERVER STARTED")

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	select {}
}