package greeting_endpoints_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/codecs/requests/hello_requests"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/codecs/responses/hello_responses"
	"github.com/split-notes/pennant-admin-backend/endpoints/greeting_endpoints"
	"github.com/split-notes/pennant-admin-backend/library/appcontext"
	"github.com/split-notes/pennant-admin-backend/mocks/services_mocks"
	"github.com/split-notes/pennant-admin-backend/services"
	"github.com/split-notes/pennant-admin-backend/services/grpc_service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHelloServer(t *testing.T) {
	var GetHelloUnitTestData = []GetHelloTestData{
		{
			BaseTestData: BaseTestData{
				Description: "happy path",
				Request:     hello_requests.GetHelloRequest{},
				Response:    hello_responses.GetHelloResponse{
					Response: []string{"one", "two"},
				},
			},
			MockGetHello: &MockGetHello{
				InRequest: grpc_service.GetHelloRequest{},
				OutReply: grpc_service.GetHelloReply{
					Greetings:  []string{"one", "two"},
				},
				OutError:     nil,
			},
		},
	}

	_ = GetHelloUnitTestData

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, data := range GetHelloUnitTestData {
		t.Run(data.Description, func(t *testing.T) {
			serviceBundle := MockGetHelloRequiredServices(mockCtrl, data)

			f := greeting_endpoints.MakeGetHelloEndpoint(appcontext.Context{}, *serviceBundle)
			requestTestData := data.Request.(hello_requests.GetHelloRequest)
			responseTestData := data.Response.(hello_responses.GetHelloResponse)
			responseData, _ := f(context.Background(), requestTestData)
			// assert results
			//assert.Equal(t, testData.ResponseCode, res.Status.Code)

			assert.Equal(t, responseTestData, responseData)
		})
	}
}

func MockGetHelloRequiredServices(mockCtrl *gomock.Controller, data GetHelloTestData) *services.Bundle {
	greeterMock := services_mocks.NewMockGreeterClient(mockCtrl)
	greeterExpect := greeterMock.EXPECT()

	if data.MockGetHello != nil {
		greeterExpect.GetHello(gomock.Any(), &data.MockGetHello.InRequest).Return(
				&data.MockGetHello.OutReply,
				data.MockGetHello.OutError)
	}

	serviceBundle := services.Bundle{
		GrpcSvc: grpc_service.Connection{
			Server:        nil,
			GreeterClient: greeterMock,
		},
	}

	return &serviceBundle
}
