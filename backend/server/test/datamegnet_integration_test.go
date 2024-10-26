package test

import (
	"datamagnet/pkg/model"
	datamagnet "datamagnet/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"server/pkg/router"
	"server/pkg/utils"
	"server/test/data"
	"testing"
)

type DataMagnetIntegrationTestSuite struct {
	suite.Suite
	engine *gin.Engine
}

func TestDataMagnetIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(DataMagnetIntegrationTestSuite))
}

func (s *DataMagnetIntegrationTestSuite) SetupSuite() {
	s.engine = router.ServerRouters(
		"/v1",
		datamagnet.IntegrationRouters,
	)
}

func (s *DataMagnetIntegrationTestSuite) Test_missing_category() {
	request := data.Request[model.AddIntegrationRequest]{}
	response := s.whenAddIntegrationFailed(request)
	s.shouldResponseStatusCode(http.StatusBadRequest, response.StatusCode())
	s.shouldResponseError("field missing: category", response.Body().Message)
}

func (s *DataMagnetIntegrationTestSuite) Test_invalid_category_format() {
	request := s.givenAddIntegrationRequest(model.AddIntegrationRequest{
		Category: "invalid category",
	})
	response := s.whenAddIntegrationFailed(request)
	s.shouldResponseStatusCode(http.StatusBadRequest, response.StatusCode())
	s.shouldResponseError("invalid format: field category only accept alphabets, digits and -", response.Body().Message)
}

func (s *DataMagnetIntegrationTestSuite) Test_preserved_category() {
	request := s.givenAddIntegrationRequest(model.AddIntegrationRequest{
		Category: "KISI",
	})
	response := s.whenAddIntegrationFailed(request)
	s.shouldResponseStatusCode(http.StatusConflict, response.StatusCode())
	s.shouldResponseError("conflict: category KISI is preserved", response.Body().Message)
}

func (s *DataMagnetIntegrationTestSuite) Test_duplicate_category() {
	request := s.givenAddIntegrationRequest(model.AddIntegrationRequest{
		Category: "duplicate-category",
	})
	_ = s.whenAddIntegrationSucceed(request)
	response := s.whenAddIntegrationFailed(request)
	s.shouldResponseStatusCode(http.StatusConflict, response.StatusCode())
	s.shouldResponseError("conflict: category duplicate-category already exists", response.Body().Message)
}

func (s *DataMagnetIntegrationTestSuite) Test_add_integration_successfully() {
	request := s.givenAddIntegrationRequest(model.AddIntegrationRequest{
		Category: "my-category",
	})
	response := s.whenAddIntegrationSucceed(request)
	s.shouldResponseStatusCode(http.StatusCreated, response.StatusCode())
	s.shouldResponseBody(
		model.AddIntegrationResponse{
			Category:   "my-category",
			WebhookUrl: "http://localhost:8080/v1/data-magnet/integrations/81192cb3-be87-4a32-be49-d038afa8d9e7",
			Secret:     "9cd72a9e3aa22efc7b615de376878de5",
			Enabled:    true,
		},
		response.Body(),
	)
}

func (s *DataMagnetIntegrationTestSuite) Test_remove_integration_successfully() {
	apitest.New().
		Handler(s.engine).
		Delete("/v1/data-magnet/integrations/81192cb3-be87-4a32-be49-d038afa8d9e7").
		Expect(s.T()).
		Status(http.StatusNoContent).
		End()
}

func (s *DataMagnetIntegrationTestSuite) givenAddIntegrationRequest(addIntegrationRequest model.AddIntegrationRequest) data.Request[model.AddIntegrationRequest] {
	return data.Request[model.AddIntegrationRequest]{
		Request: addIntegrationRequest,
	}
}

func (s *DataMagnetIntegrationTestSuite) whenAddIntegrationFailed(request data.Request[model.AddIntegrationRequest]) data.Response[model.ApiErrorResponse] {
	result := s.addIntegration(request)
	return data.Response[model.ApiErrorResponse]{Result: result}
}

func (s *DataMagnetIntegrationTestSuite) whenAddIntegrationSucceed(request data.Request[model.AddIntegrationRequest]) data.Response[model.AddIntegrationResponse] {
	result := s.addIntegration(request)
	return data.Response[model.AddIntegrationResponse]{Result: result}
}

func (s *DataMagnetIntegrationTestSuite) addIntegration(request data.Request[model.AddIntegrationRequest]) apitest.Result {
	return apitest.New().
		Handler(s.engine).
		Post("/v1/data-magnet/integrations").
		Header("Test", "true").
		Body(request.Body()).
		Expect(s.T()).
		End()
}

func (s *DataMagnetIntegrationTestSuite) shouldResponseStatusCode(expected int, actual int) bool {
	return s.Equal(expected, actual)
}

func (s *DataMagnetIntegrationTestSuite) shouldResponseError(expected string, actual string) bool {
	return s.Equal(expected, actual)
}

func (s *DataMagnetIntegrationTestSuite) shouldResponseBody(expected model.AddIntegrationResponse, actual model.AddIntegrationResponse) {
	s.True(utils.IsUuid(actual.ID))
	s.Equal(expected.Category, actual.Category)
	s.Equal(expected.WebhookUrl, actual.WebhookUrl)
	s.Equal(expected.Secret, actual.Secret)
	s.Equal(expected.Enabled, actual.Enabled)
}
