package domain

import "github.com/modelcontextprotocol/go-sdk/mcp"

type MCPTransport[In, Out any] interface {
	MCPTool() (metadata *mcp.Tool, handler mcp.ToolHandlerFor[In, Out])
}

type QueryToolInput struct {
	Query string `json:"query" jsonschema:"Query you want to perform on the Postgres database (SELECT only, maximum 50 rows)"`
}

type QueryToolOutput struct {
	Detail   string           `json:"detail" jsonschema:"Message describing the result of the query execution"`
	RowCount int              `json:"row_count" jsonschema:"Number of rows returned by the query"`
	Results  []map[string]any `json:"results" jsonschema:"Results of the query execution"`
}

type TableInfoInput struct {
	TableName string `json:"table_name" jsonschema:"Name of the table from which you want to obtain information"`
}

type TableInfoOutput struct {
	Detail  string      `json:"detail" jsonschema:"Message describing the result of the table information"`
	Results []TableInfo `json:"results" jsonschema:"Information obtained from the table"`
}

type ListTablesInput struct {
	Page   int    `json:"page" jsonschema:"Page number for paginated table results"`
	Schema string `json:"schema" jsonschema:"Database schema from which to list the tables"`
}

type ListTablesOutput struct {
	Detail   string       `json:"detail" jsonschema:"Message describing the result of the table listing"`
	RowCount int          `json:"row_count" jsonschema:"Total number of tables recovered"`
	Results  []ListTables `json:"results" jsonschema:"List of tables returned from the specified schema"`
}
