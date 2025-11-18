package tool

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/service"
)

type GetTableInfo struct {
	databaseInfoService *service.DatabaseInfoService
}

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
