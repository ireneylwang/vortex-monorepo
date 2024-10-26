package repository

import (
	"context"
	"datamagnet/internal/application/repository"
)

func NewIntegrationDao(ctx context.Context) repository.IntegrationRepository {
	inMemoryIntegrationDao := newInMemoryIntegrationDao()
	if ctx.Value("Test") == true {
		return newTestableIntegrationDao(inMemoryIntegrationDao)
	}
	return inMemoryIntegrationDao
}
