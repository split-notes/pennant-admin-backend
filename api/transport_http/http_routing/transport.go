package http_routing

import (
	"github.com/gorilla/mux"
	"github.com/split-notes/pennant-admin-backend/library/appcontext"
	"github.com/split-notes/pennant-admin-backend/services"
)

type ServerContext struct {
	AppCtx appcontext.Context
	Bundle services.Bundle
}

type Bundlable func(appCtx appcontext.Context, router *mux.Router, services services.Bundle)

var Bundlables []Bundlable

func Bundle(bundlable Bundlable) {
	Bundlables = append(Bundlables, bundlable)
}

func BundleAll(appCtx appcontext.Context, router *mux.Router, services services.Bundle) {
	for _, bundlable := range Bundlables {
		bundlable(appCtx, router, services)
	}
}
