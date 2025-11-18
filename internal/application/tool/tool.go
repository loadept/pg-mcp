package tool

import "loadept.com/pg-mcp/internal/service"

// Tool aggregates all available MCP tools for database operations.
// It provides structured access to table information, table listing, and query execution tools.
type Tool struct {
	// Tool for retrieving table structure information
	GetTableInfo *GetTableInfo
	// Tool for listing tables in a schema
	ListTables *ListTables
	// Tool for executing SQL queries
	ExecuteQuery *ExecuteQuery
}

// LoadTools creates and initializes all available MCP tools.
// It wires up each tool with its required service dependencies.
//
// Parameters:
//   - queryService: Service for executing database queries
//   - databaseInfoService: Service for retrieving database metadata
//
// Returns:
//   - *Tool: Initialized tool collection with all tools ready to use
func LoadTools(
	queryService *service.QueryService,
	databaseInfoService *service.DatabaseInfoService,
) *Tool {
	return &Tool{
		GetTableInfo: &GetTableInfo{databaseInfoService: databaseInfoService},
		ListTables:   &ListTables{databaseInfoService: databaseInfoService},
		ExecuteQuery: &ExecuteQuery{queryService: queryService},
	}
}
