package domain

import "github.com/modelcontextprotocol/go-sdk/mcp"

// MCPTransport defines the interface for MCP (Model Context Protocol) tool implementations.
// It requires a method that returns both the tool metadata and its handler function.
type MCPTransport[In, Out any] interface {
	MCPTool() (metadata *mcp.Tool, handler mcp.ToolHandlerFor[In, Out])
}

// QueryToolInput represents the input parameters for executing a database query.
type QueryToolInput struct {
	// SQL query to execute
	Query string `json:"query" jsonschema:"Query you want to perform on the Postgres database (SELECT only, maximum 50 rows)"`
}

// QueryToolOutput represents the output result of a database query execution.
type QueryToolOutput struct {
	// Execution status message
	Detail string `json:"detail" jsonschema:"Message describing the result of the query execution"`
	// Number of rows in results
	RowCount int `json:"row_count" jsonschema:"Number of rows returned by the query"`
	// Query results as key-value maps
	Results []map[string]any `json:"results" jsonschema:"Results of the query execution"`
}

// TableInfoInput represents the input parameters for retrieving table information.
type TableInfoInput struct {
	// Target table name
	TableName string `json:"table_name" jsonschema:"Name of the table from which you want to obtain information"`
}

// TableInfoOutput represents the output result of a table information request.
type TableInfoOutput struct {
	// Operation status message
	Detail string `json:"detail" jsonschema:"Message describing the result of the table information"`
	// Column information for the table
	Results []TableInfo `json:"results" jsonschema:"Information obtained from the table"`
}

// ListTablesInput represents the input parameters for listing database tables.
type ListTablesInput struct {
	// Page number for pagination (1-indexed)
	Page int `json:"page" jsonschema:"Page number for paginated table results"`
	// Schema name to query
	Schema string `json:"schema" jsonschema:"Database schema from which to list the tables"`
}

// ListTablesOutput represents the output result of a table listing request.
type ListTablesOutput struct {
	// Operation status message
	Detail string `json:"detail" jsonschema:"Message describing the result of the table listing"`
	// Number of tables returned
	RowCount int `json:"row_count" jsonschema:"Total number of tables recovered"`
	// List of table metadata
	Results []ListTables `json:"results" jsonschema:"List of tables returned from the specified schema"`
}
