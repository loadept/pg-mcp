package service

import (
	"context"

	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/repository"
)

// DatabaseInfoService provides business logic for database metadata operations.
// It acts as an intermediary between the application layer and the repository,
// handling requests for table information and table listings.
type DatabaseInfoService struct {
	repo *repository.DatabaseInfoRepository
}

// NewDatabaseInfoService creates a new instance of DatabaseInfoService.
// It initializes the service with a database info repository for metadata operations.
//
// Parameters:
//   - repo: DatabaseInfoRepository instance for querying metadata
//
// Returns:
//   - *DatabaseInfoService: Initialized service instance
func NewDatabaseInfoService(repo *repository.DatabaseInfoRepository) *DatabaseInfoService {
	return &DatabaseInfoService{repo: repo}
}

// GetTableInfo retrieves detailed structural information about a specific table.
// It delegates the request to the repository layer without additional processing.
//
// Parameters:
//   - ctx: Context for operation cancellation and timeout control
//   - tableName: Name of the table to retrieve information for
//
// Returns:
//   - []domain.TableInfo: Slice containing column information for the table
//   - error: Any error encountered during the operation
func (s *DatabaseInfoService) GetTableInfo(
	ctx context.Context,
	tableName string,
) ([]domain.TableInfo, error) {
	return s.repo.GetTableInfo(ctx, tableName)
}

// ListTables retrieves a paginated list of tables from a specified schema.
// It delegates the request to the repository layer without additional processing.
//
// Parameters:
//   - ctx: Context for operation cancellation and timeout control
//   - page: Page number for pagination (1-indexed)
//   - schema: Database schema name to list tables from
//
// Returns:
//   - []domain.ListTables: Slice containing table metadata
//   - error: Any error encountered during the operation
func (s *DatabaseInfoService) ListTables(
	ctx context.Context,
	page int,
	schema string,
) ([]domain.ListTables, error) {
	return s.repo.ListTables(ctx, page, schema)
}
