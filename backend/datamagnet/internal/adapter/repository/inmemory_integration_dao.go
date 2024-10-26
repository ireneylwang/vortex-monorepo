package repository

import (
	"datamagnet/internal/domain/entity"
	"github.com/google/uuid"
)

var (
	IntegrationRecordList []*entity.Integration
)

type InMemoryIntegrationDao struct {
}

func newInMemoryIntegrationDao() *InMemoryIntegrationDao {
	return &InMemoryIntegrationDao{}
}

func (d *InMemoryIntegrationDao) ExistsByOrgIdAndCategory(orgId string, category entity.Category) bool {
	for _, integration := range IntegrationRecordList {
		if integration.OrgID == orgId && integration.Category == category {
			return true
		}
	}

	return false
}

func (d *InMemoryIntegrationDao) Save(integration *entity.Integration) *entity.Integration {
	integration.ID = uuid.NewString()
	IntegrationRecordList = append(IntegrationRecordList, integration)
	return integration
}
