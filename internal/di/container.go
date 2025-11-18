package di

import (
	"database/sql"

	"loadept.com/pg-mcp/internal/repository"
	"loadept.com/pg-mcp/internal/service"
)

type Container struct {
	DatabaseInfoRepository *repository.DatabaseInfoRepository
	QueryRepository        *repository.QueryRepository

	DatabaseInfoService *service.DatabaseInfoService
	QueryService        *service.QueryService
}

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
