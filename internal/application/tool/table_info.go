package tool

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/service"
)

// GetTableInfo implements an MCP tool for retrieving table structure information.
// It provides detailed metadata about table columns including names, data types,
// and nullability constraints.
type GetTableInfo struct {
	databaseInfoService *service.DatabaseInfoService
}

// MCPTool returns the metadata and handler for the get_table_info MCP tool.
// The tool retrieves structural information about a specified database table.
//
// Returns:
//   - *mcp.Tool: Tool metadata including name and description
//   - mcp.ToolHandlerFor: Handler function that processes table info requests
func (m *GetTableInfo) MCPTool() (
	*mcp.Tool,
	mcp.ToolHandlerFor[domain.TableInfoInput, domain.TableInfoOutput],
) {
	metadata := &mcp.Tool{
		Name:        "get_table_info",
		Description: "Get information about a table",
	}

	handler := func(
		ctx context.Context,
		req *mcp.CallToolRequest,
		input domain.TableInfoInput,
	) (
		*mcp.CallToolResult,
		domain.TableInfoOutput,
		error,
	) {
		results, err := m.databaseInfoService.GetTableInfo(ctx, input.TableName)
		if err != nil {
			output := domain.TableInfoOutput{
				Detail:  fmt.Sprintf("An error occurred while retrieving table information: %v", err.Error()),
				Results: []domain.TableInfo{},
			}
			return nil, output, err
		}

		output := domain.TableInfoOutput{
			Detail:  "The table information has been retrieved successfully.",
			Results: results,
		}
		return nil, output, nil
	}

	return metadata, handler
}
