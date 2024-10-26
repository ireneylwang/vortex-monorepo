package model

import (
	"datamagnet/internal/adapter/controller"
	"datamagnet/internal/adapter/presenter"
	err "server/pkg/presenter"
)

type ApiErrorResponse err.ApiErrorViewModel

type AddIntegrationRequest controller.AddIntegrationRequest

type AddIntegrationResponse presenter.AddIntegrationViewModel
