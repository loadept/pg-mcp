package service

import (
	"context"
	"fmt"

	"loadept.com/pg-mcp/internal/repository"
)

type QueryService struct {
	repo *repository.QueryRepository
}

func NewQueryService(repo *repository.QueryRepository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) ExecuteQuery(
	ctx context.Context,
	query string,
	args ...any,
) ([]map[string]any, error) {
	if query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}

	return s.repo.ExecuteQuery(ctx, query, args...)
}
