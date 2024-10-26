package presenter

import (
	"datamagnet/internal/application/usecase"
	"net/http"
)

type AddIntegrationViewModel struct {
	ID         string
	Category   string
	WebhookUrl string
	Secret     string
	Enabled    bool
}

type AddIntegrationPresenter struct {
	viewModel AddIntegrationViewModel
}

func (p *AddIntegrationPresenter) Present(output *usecase.AddIntegrationOutput) {
	p.viewModel = AddIntegrationViewModel{
		ID:         output.ID,
		Category:   output.Category,
		WebhookUrl: output.WebhookUrl,
		Secret:     output.Secret,
		Enabled:    output.Enabled,
	}
}

func (p *AddIntegrationPresenter) StatusCode() int {
	return http.StatusCreated
}

func (p *AddIntegrationPresenter) ApiDataResponse() AddIntegrationViewModel {
	return p.viewModel
}
