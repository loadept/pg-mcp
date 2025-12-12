package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/application"
	"loadept.com/pg-mcp/internal/application/tool"
	"loadept.com/pg-mcp/internal/config"
	"loadept.com/pg-mcp/internal/di"
)

var VERSION = "dev"

var version = flag.Bool("version", false, "Show the application version")

func init() {
	config.LoadEnvs()
}

func main() {
	uri := flag.String("u", "", "PostgreSQL connection URI")
	flag.Parse()

	if *version {
		fmt.Printf("pg-mcp version %s\n", VERSION)
		os.Exit(0)
	}

	pg, err := config.NewDBPostgres(*uri)
	if err != nil {
		fmt.Println("An error occurred while connecting to the database:", err)
		os.Exit(1)
	}
	defer pg.Close()

	implementation := &mcp.Implementation{
		Name:    "PostgreSQL MCP Server",
		Version: VERSION,
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
