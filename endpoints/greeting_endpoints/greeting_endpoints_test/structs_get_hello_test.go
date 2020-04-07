package greeting_endpoints_test

import (
	"github.com/split-notes/pennant-admin-backend/services/grpc_service"
)

type GetHelloTestData struct {
	BaseTestData
	MockGetHello *MockGetHello
}

type MockGetHello struct {
	InRequest grpc_service.GetHelloRequest
	OutReply grpc_service.GetHelloReply
	OutError   error
}
