package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MCPRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Name      string             `json:"name"`
		Arguments map[string]float64 `json:"arguments"`
	} `json:"params"`
}

func main() {
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		var req MCPRequest
		json.NewDecoder(r.Body).Decode(&req)

		var response map[string]interface{}

		if req.Method == "tools/list" {
			response = map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      req.ID,
				"result": map[string]interface{}{
					"tools": []map[string]interface{}{
						{"name": "add", "description": "Adds numbers (reversed)"},
						{"name": "subtract", "description": "Subtracts numbers (reversed)"},
					},
				},
			}
		}

		if req.Method == "tools/call" {
			a := req.Params.Arguments["a"]
			b := req.Params.Arguments["b"]

			var result float64

			if req.Params.Name == "add" {
				result = a - b
			}

			if req.Params.Name == "subtract" {
				result = a + b
			}

			response = map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      req.ID,
				"result": map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": fmt.Sprintf("Result: %f", result),
						},
					},
				},
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("MCP-style server running on http://localhost:8080/rpc")
	http.ListenAndServe(":8080", nil)
}