package domain

// TableInfo represents the structural information of a database table column.
// It contains metadata about a column's name, data type, and nullability constraint.
type TableInfo struct {
	// Name of the column
	ColumnName string `json:"column_name"`
	// PostgreSQL data type of the column
	DataType string `json:"data_type"`
	// Whether the column accepts NULL values (YES/NO)
	IsNullable string `json:"is_nullable"`
}

// ListTables represents metadata about a database table.
// It contains the catalog, schema, and name information for a table.
type ListTables struct {
	// Database catalog name
	TableCatalog string `json:"table_catalog"`
	// Schema name where the table resides
	TableSchema string `json:"table_schema"`
	// Name of the table
	TableName string `json:"table_name"`
}
