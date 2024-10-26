package usecase

import (
	"datamagnet/internal/application/repository"
	"datamagnet/internal/domain/entity"
	"fmt"
	"server/pkg/model"
	"server/pkg/utils"
)

type AddIntegrationInput struct {
	OrgID    string
	Category entity.Category
}

type AddIntegrationInputOption func(*AddIntegrationInput)

func NewAddIntegrationInput(options ...AddIntegrationInputOption) *AddIntegrationInput {
	input := new(AddIntegrationInput)
	for _, option := range options {
		option(input)
	}
	return input
}

func WithCategory(category string) AddIntegrationInputOption {
	return func(input *AddIntegrationInput) {
		input.Category = entity.Category(category)
	}
}

type AddIntegrationOutput struct {
	ID         string
	OrgID      string
	Category   string
	WebhookUrl string
	Secret     string
	Enabled    bool
}

func NewAddIntegrationOutput(integration *entity.Integration) *AddIntegrationOutput {
	return &AddIntegrationOutput{
		ID:         integration.ID,
		Category:   integration.Category.String(),
		WebhookUrl: integration.WebhookUrl(),
		Secret:     integration.Secret,
		Enabled:    integration.Enabled,
	}
}

type AddIntegrationUseCase struct {
	integrationRepository repository.IntegrationRepository
}

func NewAddIntegrationUseCase(integrationRepository repository.IntegrationRepository) *AddIntegrationUseCase {
	return &AddIntegrationUseCase{
		integrationRepository: integrationRepository,
	}
}

func (c *AddIntegrationUseCase) Execute(input *AddIntegrationInput) (*AddIntegrationOutput, error) {
	if utils.ValidateMissingField(input.Category.String()) {
		return nil, model.NewFieldMissingError("category")
	}

	if utils.ValidateInvalidFormat(`^[a-zA-Z0-9-]+$`, input.Category.String()) {
		return nil, model.NewFieldInvalidFormatError("field category only accept alphabets, digits and -")
	}

	if input.Category.IsPreserved() {
		return nil, model.NewConflictError(fmt.Sprintf("category %s is preserved", input.Category))
	}

	if exists := c.integrationRepository.ExistsByOrgIdAndCategory(input.OrgID, input.Category); exists {
		return nil, model.NewConflictError(fmt.Sprintf("category %s already exists", input.Category))
	}

	integration := entity.NewIntegration(input.OrgID, input.Category)
	integration = c.integrationRepository.Save(integration)
	return NewAddIntegrationOutput(integration), nil
}
