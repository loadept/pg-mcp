package application

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
)

func LoadTool[In, Out any](server *mcp.Server, tool domain.MCPTransport[In, Out]) {
	metadata, handler := tool.MCPTool()
	mcp.AddTool(server, metadata, handler)
}
