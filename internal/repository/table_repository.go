package repository

import (
	"context"
	"database/sql"

	"loadept.com/pg-mcp/internal/domain"
)

type DatabaseInfoRepository struct {
	db *sql.DB
}

func NewDatabaseInfoRepository(db *sql.DB) *DatabaseInfoRepository {
	return &DatabaseInfoRepository{db: db}
}

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
