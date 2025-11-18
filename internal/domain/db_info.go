package domain

type TableInfo struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
	IsNullable string `json:"is_nullable"`
}

type ListTables struct {
	TableCatalog string `json:"table_catalog"`
	TableSchema  string `json:"table_schema"`
	TableName    string `json:"table_name"`
}
