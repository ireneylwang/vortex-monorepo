package entity

import (
	"server/pkg/env"
)

type Integration struct {
	ID       string
	OrgID    string
	Category Category
	Secret   string
	Enabled  bool
}

func NewIntegration(OrgID string, category Category) *Integration {
	return &Integration{
		OrgID:    OrgID,
		Category: category,
		Secret:   hash(),
		Enabled:  true,
	}
}

func (i *Integration) WebhookUrl() string {
	return env.OpenApiHost() + "/v1/data-magnet/integrations/" + i.ID
}

func hash() string {
	return "9cd72a9e3aa22efc7b615de376878de5"
}
