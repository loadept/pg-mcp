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
	db := pg.GetDB()

	implementation := &mcp.Implementation{
		Name:    "PostgreSQL MCP Server",
		Version: "0.1.0",
	}
	server := mcp.NewServer(implementation, nil)

	containerDependencies := di.NewContainer(db)
	tool := tool.GetTools(
		containerDependencies.QueryService,
		containerDependencies.DatabaseInfoService,
	)

	application.LoadTool(server, tool.GetTableInfo)
	application.LoadTool(server, tool.ListTables)
	application.LoadTool(server, tool.ExecuteQuery)

	log.Println("MCP server is running...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
