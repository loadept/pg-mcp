package di

import (
	"database/sql"

	"loadept.com/pg-mcp/internal/repository"
	"loadept.com/pg-mcp/internal/service"
)

// Container holds all application dependencies for dependency injection.
// It provides centralized access to repositories and services, ensuring
// consistent wiring of dependencies throughout the application.
type Container struct {
	// Repository for database metadata operations
	DatabaseInfoRepository *repository.DatabaseInfoRepository
	// Repository for query execution
	QueryRepository *repository.QueryRepository

	// Service for database metadata operations
	DatabaseInfoService *service.DatabaseInfoService
	// Service for query execution
	QueryService *service.QueryService
}

// NewContainer creates and initializes a new dependency injection container.
// It wires up all repositories and services with their required dependencies,
// following the dependency injection pattern.
//
// Parameters:
//   - db: Active database connection pool to be used by all repositories
//
// Returns:
//   - *Container: Fully initialized container with all dependencies wired up
func NewContainer(db *sql.DB) *Container {
	dbInfoRepo := repository.NewDatabaseInfoRepository(db)
	queryRepo := repository.NewQueryRepository(db)

	dbInfoService := service.NewDatabaseInfoService(dbInfoRepo)
	queryService := service.NewQueryService(queryRepo)

	return &Container{
		DatabaseInfoRepository: dbInfoRepo,
		QueryRepository:        queryRepo,
		DatabaseInfoService:    dbInfoService,
		QueryService:           queryService,
	}
}
