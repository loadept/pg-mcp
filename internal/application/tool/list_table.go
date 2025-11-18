package tool

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/service"
)

type ListTables struct {
	databaseInfoService *service.DatabaseInfoService
}

func (m *ListTables) MCPTool() (
	*mcp.Tool,
	mcp.ToolHandlerFor[domain.ListTablesInput, domain.ListTablesOutput],
) {
	metadata := &mcp.Tool{
		Name:        "list_tables",
		Description: "List all available tables in a schema",
	}

	handler := func(
		ctx context.Context,
		req *mcp.CallToolRequest,
		input domain.ListTablesInput,
	) (
		*mcp.CallToolResult,
		domain.ListTablesOutput,
		error,
	) {
		results, err := m.databaseInfoService.ListTables(ctx, input.Page, input.Schema)
		if err != nil {
			output := domain.ListTablesOutput{
				Detail:  fmt.Sprintf("An error occurred while listing tables: %v", err.Error()),
				Results: []domain.ListTables{},
			}
			return nil, output, err
		}

		output := domain.ListTablesOutput{
			Detail:  "The table listing has been executed successfully.",
			Results: results,
		}
		return nil, output, nil
	}

	return metadata, handler
}
