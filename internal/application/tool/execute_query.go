package tool

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/service"
)

type ExecuteQuery struct {
	queryService *service.QueryService
}

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
