package controller

import (
	"datamagnet/internal/adapter/presenter"
	"datamagnet/internal/adapter/repository"
	"datamagnet/internal/application/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	apierr "server/pkg/presenter"
)

type AddIntegrationRequest struct {
	Category string `json:"category"`
}

func NewAddIntegrationRequest(ctx *gin.Context) *AddIntegrationRequest {
	request := new(AddIntegrationRequest)
	if err := ctx.BindJSON(request); err != nil {
		fmt.Println(err)
	}
	return request
}

func (r *AddIntegrationRequest) Input() *usecase.AddIntegrationInput {
	return usecase.NewAddIntegrationInput(
		usecase.WithCategory(r.Category),
	)
}

func AddIntegration(ctx *gin.Context) {
	ctx.Set("Test", ctx.Request.Header["Test"][0] == "true")

	dao := repository.NewIntegrationDao(ctx)
	useCase := usecase.NewAddIntegrationUseCase(dao)

	input := NewAddIntegrationRequest(ctx).Input()
	output, err := useCase.Execute(input)

	if err != nil {
		errorPresenter := new(apierr.ApiErrorPresenter)
		errorPresenter.Present(err)
		ctx.JSON(errorPresenter.StatusCode(), errorPresenter.ApiErrorResponse())
		return
	}

	integrationPresenter := new(presenter.AddIntegrationPresenter)
	integrationPresenter.Present(output)
	ctx.JSON(integrationPresenter.StatusCode(), integrationPresenter.ApiDataResponse())
}
