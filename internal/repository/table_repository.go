package repository

import (
	"context"
	"database/sql"

	"loadept.com/pg-mcp/internal/domain"
)

// DatabaseInfoRepository handles database metadata operations.
// It provides methods to query information about database tables and their structures
// using PostgreSQL's information_schema views.
type DatabaseInfoRepository struct {
	db *sql.DB
}

// NewDatabaseInfoRepository creates a new instance of DatabaseInfoRepository.
// It initializes the repository with a database connection that will be used
// for all metadata queries.
//
// Parameters:
//   - db: Active database connection pool
//
// Returns:
//   - *DatabaseInfoRepository: Initialized repository instance
func NewDatabaseInfoRepository(db *sql.DB) *DatabaseInfoRepository {
	return &DatabaseInfoRepository{db: db}
}

// GetTableInfo retrieves detailed information about a specific table's structure.
// It queries the information_schema.columns view to obtain column names, data types,
// and nullability constraints for the specified table.
//
// Parameters:
//   - ctx: Context for query cancellation and timeout control
//   - tableName: Name of the table to retrieve information for
//
// Returns:
//   - []domain.TableInfo: Slice containing column information for each column in the table
//   - error: Any error encountered during the query execution
func (r *DatabaseInfoRepository) GetTableInfo(
	ctx context.Context,
	tableName string,
) ([]domain.TableInfo, error) {
	query := `SELECT
		column_name
		, data_type
		, is_nullable
	FROM information_schema.columns
	WHERE table_name = $1`

	rows, err := r.db.QueryContext(ctx, query, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tableInfo []domain.TableInfo
	for rows.Next() {
		var table domain.TableInfo
		if err := rows.Scan(
			&table.ColumnName,
			&table.DataType,
			&table.IsNullable,
		); err != nil {
			return nil, err
		}
		tableInfo = append(tableInfo, table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if tableInfo == nil {
		tableInfo = []domain.TableInfo{}
	}
	return tableInfo, nil
}

// ListTables retrieves a paginated list of tables from a specific database schema.
// It queries the information_schema.tables view to obtain table metadata,
// returning only base tables (excluding views) in descending order by table name.
//
// Parameters:
//   - ctx: Context for query cancellation and timeout control
//   - page: Page number for pagination (1-indexed), each page contains 10 tables
//   - schema: Database schema name to list tables from (e.g., "public")
//
// Returns:
//   - []domain.ListTables: Slice containing table catalog, schema, and name information
//   - error: Any error encountered during the query execution
func (r *DatabaseInfoRepository) ListTables(
	ctx context.Context,
	page int,
	schema string,
) ([]domain.ListTables, error) {
	query := `SELECT
		table_catalog
		, table_schema
		, table_name
	FROM information_schema.tables
	WHERE table_type = 'BASE TABLE'
	AND table_schema = $1
	ORDER BY 3 DESC
	LIMIT 10
	OFFSET $2`

	offset := (page - 1) * 10
	rows, err := r.db.QueryContext(ctx, query, schema, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []domain.ListTables
	for rows.Next() {
		var table domain.ListTables
		if err := rows.Scan(
			&table.TableCatalog,
			&table.TableSchema,
			&table.TableName,
		); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if tables == nil {
		tables = []domain.ListTables{}
	}
	return tables, nil
}
