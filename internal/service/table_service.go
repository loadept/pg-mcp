package service

import (
	"context"

	"loadept.com/pg-mcp/internal/domain"
	"loadept.com/pg-mcp/internal/repository"
)

type DatabaseInfoService struct {
	repo *repository.DatabaseInfoRepository
}

func NewDatabaseInfoService(repo *repository.DatabaseInfoRepository) *DatabaseInfoService {
	return &DatabaseInfoService{repo: repo}
}

func (s *DatabaseInfoService) GetTableInfo(
	ctx context.Context,
	tableName string,
) ([]domain.TableInfo, error) {
	return s.repo.GetTableInfo(ctx, tableName)
}

func (s *DatabaseInfoService) ListTables(
	ctx context.Context,
	page int,
	schema string,
) ([]domain.ListTables, error) {
	return s.repo.ListTables(ctx, page, schema)
}
