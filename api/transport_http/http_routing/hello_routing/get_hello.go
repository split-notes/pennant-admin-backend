package hello_routing

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/codecs/requests/hello_requests"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/codecs/responses/hello_responses"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/http_routing"
	"github.com/split-notes/pennant-admin-backend/endpoints/greeting_endpoints"
	"github.com/split-notes/pennant-admin-backend/library/appcontext"
	"github.com/split-notes/pennant-admin-backend/services"
)

func init() {
	http_routing.Bundle(MakeGetHelloHttpHandler())
}

func MakeGetHelloHttpHandler() http_routing.Bundlable {
	return func(appCtx appcontext.Context, router *mux.Router, services services.Bundle) {
		api := router.PathPrefix("/api").Subrouter()

		endpoint := greeting_endpoints.MakeGetHelloEndpoint(appCtx, services)
		decoder, _ := hello_requests.MakeGetHelloRequestDecoder(appCtx)
		encoder, _ := hello_responses.MakeGetHelloResponseEncoder(appCtx)

		//POST /Find Campaing ID
		api.Methods("GET").Path("/hello").Handler(httpTransport.NewServer(
			endpoint,
			decoder,
			encoder,
		))
	}
}
