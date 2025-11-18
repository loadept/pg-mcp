package repository

import (
	"context"
	"database/sql"
)

type QueryRepository struct {
	db *sql.DB
}

func NewQueryRepository(db *sql.DB) *QueryRepository {
	return &QueryRepository{db: db}
}

func (r *QueryRepository) ExecuteQuery(
	ctx context.Context,
	query string,
	args ...any,
) ([]map[string]any, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly:  true,
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	for rows.Next() {
		values := make([]any, len(columns))
		valuePointers := make([]any, len(columns))
		for i := range columns {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		rowData := make(map[string]any)
		for i, columnName := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowData[columnName] = string(b)
			} else {
				rowData[columnName] = val
			}
		}
		results = append(results, rowData)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	if results == nil {
		results = []map[string]any{}
	}
	return results, nil
}
