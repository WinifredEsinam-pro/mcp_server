package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Request struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      *int            `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ToolCallParams struct {
	Name  string                 `json:"name"`
	Input map[string]interface{} `json:"input"`
}

func main() {
	log("MCP server started")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var req Request

		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
			log("Invalid JSON:", err)
			continue
		}
		if req.ID == nil {
			handleNotification(req)
			continue
		}

		handleRequest(req)
	}
}

func handleRequest(req Request) {
	switch req.Method {

	case "initialize":
		sendResult(*req.ID, map[string]interface{}{
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]string{
				"name":    "go-mcp-server",
				"version": "1.0.0",
			},
		})

	case "tools/list":
		sendResult(*req.ID, map[string]interface{}{
			"tools": []map[string]interface{}{
				{
					"name":        "add",
					"description": "Adds numbers (reversed logic)",
					"inputSchema": schema(),
				},
				{
					"name":        "subtract",
					"description": "Subtracts numbers (reversed logic)",
					"inputSchema": schema(),
				},
			},
		})

	case "tools/call":
		handleToolCall(req)

	default:
		sendError(*req.ID, -32601, "Method not found")
	}
}

func handleToolCall(req Request) {
	var params ToolCallParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		sendError(*req.ID, -32602, "Invalid params")
		return
	}

	a, okA := params.Input["a"].(float64)
	b, okB := params.Input["b"].(float64)

	if !okA || !okB {
		sendError(*req.ID, -32602, "Invalid input values")
		return
	}

	var result float64

	switch params.Name {
	case "add":
		result = a - b 
	case "subtract":
		result = a + b
	default:
		sendError(*req.ID, -32601, "Unknown tool")
		return
	}

	sendResult(*req.ID, map[string]interface{}{
		"content": []map[string]interface{}{
			{
				"type": "text",
				"text": fmt.Sprintf("Result: %f", result),
			},
		},
	})
}


func handleNotification(req Request) {
	switch req.Method {
	case "initialized":
		log("Client initialized connection")
	default:
		log("Unknown notification:", req.Method)
	}
}

func sendResult(id int, result interface{}) {
	res := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  result,
	}
	json.NewEncoder(os.Stdout).Encode(res)
}

func sendError(id int, code int, message string) {
	res := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: message,
		},
	}
	json.NewEncoder(os.Stdout).Encode(res)
}

func log(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func schema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"a": map[string]string{"type": "number"},
			"b": map[string]string{"type": "number"},
		},
		"required": []string{"a", "b"},
	}
}