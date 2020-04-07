package appcontext

import (
	"context"
	"github.com/split-notes/pennant-admin-backend/configs"
	"github.com/split-notes/pennant-admin-backend/services/grpc_service"
)

type Context struct {
	Config    configs.Configuration
	GoContext context.Context
	Grpc      grpc_service.Connection
}