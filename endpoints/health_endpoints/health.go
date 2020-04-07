package health_endpoints

import "net/http"

// liveliness probe.
func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}