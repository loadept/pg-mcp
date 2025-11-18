package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/application"
	"loadept.com/pg-mcp/internal/application/tool"
	"loadept.com/pg-mcp/internal/config"
	"loadept.com/pg-mcp/internal/di"
)

func init() {
	config.LoadEnvs()
}

func main() {
	pg, err := config.NewDBPostgres()
	if err != nil {
		fmt.Println("An error occurred while connecting to the database:", err)
		os.Exit(1)
	}
	defer pg.Close()

	implementation := &mcp.Implementation{
		Name:    "PostgreSQL MCP Server",
		Version: "0.2.0",
	}
	server := mcp.NewServer(implementation, nil)

	containerDependencies := di.NewContainer(pg.GetDB())
	tool := tool.LoadTools(
		containerDependencies.QueryService,
		containerDependencies.DatabaseInfoService,
	)

	application.AddTool(server, tool.GetTableInfo)
	application.AddTool(server, tool.ListTables)
	application.AddTool(server, tool.ExecuteQuery)

	log.Println("MCP server is running...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
