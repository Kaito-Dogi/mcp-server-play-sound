// Package main provides a simple example MCP server with a greet tool.
// This is a "Hello World" example demonstrating basic MCP tool registration.
package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

func SayHello(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input Input,
) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	return nil, Output{Greeting: "Hello, " + input.Name + "!"}, nil
}

func main() {
	log.Println("example-greet: starting...")

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{Name: "example-greet", Version: "v0.0.1"}, nil)
	log.Println("example-greet: created server")

	// Register tools
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Say Hello"}, SayHello)
	log.Println("example-greet: registered tool: greet")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}

	log.Println("example-greet: server stopped")
}
