package service

import (
	"context"
	"fmt"

	"loadept.com/pg-mcp/internal/repository"
)

// QueryService provides business logic for executing database queries.
// It acts as an intermediary between the application layer and the repository,
// adding validation and additional processing as needed.
type QueryService struct {
	repo *repository.QueryRepository
}

// NewQueryService creates a new instance of QueryService.
// It initializes the service with a query repository for database operations.
//
// Parameters:
//   - repo: QueryRepository instance for executing queries
//
// Returns:
//   - *QueryService: Initialized service instance
func NewQueryService(repo *repository.QueryRepository) *QueryService {
	return &QueryService{repo: repo}
}

// ExecuteQuery validates and executes a SQL query through the repository layer.
// It performs basic validation to ensure the query is not empty before execution.
//
// Parameters:
//   - ctx: Context for query cancellation and timeout control
//   - query: SQL query string to execute
//
// Returns:
//   - []map[string]any: Query results as a slice of maps
//   - error: Validation or execution error, if any
func (s *QueryService) ExecuteQuery(
	ctx context.Context,
	query string,
) ([]map[string]any, error) {
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}

	return s.repo.ExecuteQuery(ctx, query)
}
