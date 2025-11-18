package tool

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/service"
)

// ExecuteQuery implements an MCP tool for executing SQL queries on the database.
// It provides a safe interface for running SELECT queries with automatic
// error handling and result formatting.
type ExecuteQuery struct {
	queryService *service.QueryService
}

// MCPTool returns the metadata and handler for the execute_query MCP tool.
// The tool allows executing SELECT queries on the PostgreSQL database with
// a maximum result limit of 50 rows.
//
// Returns:
//   - *mcp.Tool: Tool metadata including name and description
//   - mcp.ToolHandlerFor: Handler function that processes query requests
func (m *ExecuteQuery) MCPTool() (
	*mcp.Tool,
	mcp.ToolHandlerFor[domain.QueryToolInput, domain.QueryToolOutput],
) {
	metadata := &mcp.Tool{
		Name:        "execute_query",
		Description: "Executes a query on the postgres database",
	}

	handler := func(ctx context.Context, req *mcp.CallToolRequest, input domain.QueryToolInput) (
		*mcp.CallToolResult,
		domain.QueryToolOutput,
		error,
	) {
		results, err := m.queryService.ExecuteQuery(ctx, input.Query)
		if err != nil {
			output := domain.QueryToolOutput{
				Detail:   fmt.Sprintf("An error occurred while executing the query: %v", err.Error()),
				RowCount: 0,
				Results:  []map[string]any{},
			}
			return nil, output, err
		}

		output := domain.QueryToolOutput{
			Detail:   "Query executed successfully.",
			RowCount: len(results),
			Results:  results,
		}
		return nil, output, nil
	}

	return metadata, handler
}
