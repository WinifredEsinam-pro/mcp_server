package main
import (
	"fmt"
	"bufio"
	"encoding/json"
	"os"
)

type Request struct{
	Jsonrpc string `json:"jsonrpc"`
	ID int `json:"id, omitempty"`
	Method string `json:"method"`
    Params json.RawMessage `json:"params,omitempty"`
}

type Response struct{
	Jsonrpc string `json:"jsonrpc"`
	ID int `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
}

type RPCError struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

type ToolCallParams struct{
	Name string `json:"name"`
	Input map[string]interface{} `json:"input"`
}

func main(){
	log("MCP server started")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		var req Request
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil{
			log("Invalid JSON:", err)
			continue
		}
		if req.ID == nil{
			handleNotification(req)
			continue
		}
		handleNotification(req)
	}
}

func handleRequest(req Request){
	switch req.Method{
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