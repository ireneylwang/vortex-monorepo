package presenter

import (
	"errors"
	"fmt"
	"server/pkg/model"
)

type ApiErrorViewModel struct {
	Message string `json:"message"`
}

type ApiErrorPresenter struct {
	statusCode int
	viewModel  ApiErrorViewModel
}

func (p *ApiErrorPresenter) Present(err error) {
	apiError := new(model.ApiError)
	if !errors.As(err, &apiError) {
		fmt.Printf("error: %+v\n", err)
		apiError = model.NewInternalServerError()
	}

	p.statusCode = apiError.StatusCode()
	p.viewModel = ApiErrorViewModel{
		Message: apiError.Error(),
	}
}

func (p *ApiErrorPresenter) StatusCode() int {
	return p.statusCode
}

func (p *ApiErrorPresenter) ApiErrorResponse() ApiErrorViewModel {
	return p.viewModel
}
