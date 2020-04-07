package health_endpoints

import "net/http"

// readiness probe.
func ready(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
