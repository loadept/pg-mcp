package tool

import "loadept.com/pg-mcp/internal/service"

type Tool struct {
	GetTableInfo *GetTableInfo
	ListTables   *ListTables
	ExecuteQuery *ExecuteQuery
}

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
