package repository

import "datamagnet/internal/domain/entity"

type IntegrationRepository interface {
	ExistsByOrgIdAndCategory(orgId string, category entity.Category) bool
	Save(integration *entity.Integration) *entity.Integration
}
