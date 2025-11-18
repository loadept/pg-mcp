package application

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"loadept.com/pg-mcp/internal/domain"
)

// AddTool registers an MCP tool with the given server.
// It extracts the tool's metadata and handler from the MCPTransport implementation
// and registers them with the MCP server.
//
// Type Parameters:
//   - In: Input type for the tool
//   - Out: Output type for the tool
//
// Parameters:
//   - server: MCP server instance to register the tool with
//   - tool: MCPTransport implementation containing the tool logic
func AddTool[In, Out any](server *mcp.Server, tool domain.MCPTransport[In, Out]) {
	metadata, handler := tool.MCPTool()
	mcp.AddTool(server, metadata, handler)
}
