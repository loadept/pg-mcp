package repository

import (
	"context"
	"database/sql"
)

// QueryRepository handles the execution of read-only database queries.
// It provides methods to safely execute SELECT queries with transaction management
// and automatic result marshaling.
type QueryRepository struct {
	db *sql.DB
}

// NewQueryRepository creates a new instance of QueryRepository.
// It initializes the repository with a database connection that will be used
// for executing read-only queries.
//
// Parameters:
//   - db: Active database connection pool
//
// Returns:
//   - *QueryRepository: Initialized repository instance
func NewQueryRepository(db *sql.DB) *QueryRepository {
	return &QueryRepository{db: db}
}

// ExecuteQuery executes a read-only SQL query and returns the results as a slice of maps.
// The method runs the query within a read-only transaction at READ COMMITTED isolation level.
// Results are automatically converted to a generic map structure with column names as keys.
// Byte arrays are converted to strings for easier handling.
//
// Parameters:
//   - ctx: Context for query cancellation and timeout control
//   - query: SQL query to execute (should be a SELECT statement)
//
// Returns:
//   - []map[string]any: Slice of maps where each map represents a row with column names as keys
//   - error: Any error encountered during query execution or result processing
func (r *QueryRepository) ExecuteQuery(
	ctx context.Context,
	query string,
) ([]map[string]any, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly:  true,
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get column names from the result set
	// return slice of strings e.g. ["id", "name", "age"]
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	for rows.Next() {
		// Create slices to hold column values and their pointers
		values := make([]any, len(columns))
		valuePointers := make([]any, len(columns))
		for i := range columns {
			// Assign pointer to each value in slice
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		// Map column values to their respective column names
		rowData := make(map[string]any)
		for i, columnName := range columns {
			val := values[i]
			// Convert byte arrays to strings (this is common for text/varchar columns)
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
