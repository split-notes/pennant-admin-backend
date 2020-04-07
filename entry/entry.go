package entry

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/split-notes/pennant-admin-backend/api/transport_http/http_routing"
	"github.com/split-notes/pennant-admin-backend/configs"
	"github.com/split-notes/pennant-admin-backend/library/appcontext"
	"github.com/split-notes/pennant-admin-backend/services"
	"github.com/split-notes/pennant-admin-backend/services/grpc_service"
	"log"
	"net/http"
)

func Entry() {
	// Get Configs
	config, err := configs.Configure()
	if err != nil {
		panic(err) }

	grpcConnection, err := grpc_service.Start(*config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := grpc_service.Stop(); err != nil {
			panic(err)
		}
	}()

	// Build Context
	ctx := appcontext.Context{
		Config: *config,
		Grpc: *grpcConnection,
		// Redis
	}

	// Bundle Services
	serviceBundle, err := services.NewBundle(ctx)
	if err != nil {
		panic(err) }

	serviceBundle.GrpcSvc = *grpcConnection

	router := mux.NewRouter()
	http_routing.BundleAll(ctx, router, *serviceBundle)

	svr := http.Server{
		Addr:    fmt.Sprintf(":%s", config.SrvPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			router.ServeHTTP(w, r)
		}),
	}

	log.Println("Starting Server...")
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatalf("error listening: %v", err.Error())
	}
}
