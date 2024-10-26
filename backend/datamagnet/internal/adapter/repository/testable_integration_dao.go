package repository

import (
	"datamagnet/internal/application/repository"
	"datamagnet/internal/domain/entity"
)

type TestableIntegrationDao struct {
	dao repository.IntegrationRepository
}

func newTestableIntegrationDao(dao repository.IntegrationRepository) *TestableIntegrationDao {
	return &TestableIntegrationDao{dao: dao}
}

func (d *TestableIntegrationDao) ExistsByOrgIdAndCategory(orgId string, category entity.Category) bool {
	return d.dao.ExistsByOrgIdAndCategory(orgId, category)
}

func (d *TestableIntegrationDao) Save(integration *entity.Integration) *entity.Integration {
	save := d.dao.Save(integration)
	save.ID = "81192cb3-be87-4a32-be49-d038afa8d9e7"
	return save
}
