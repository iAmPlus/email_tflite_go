package app

import (
	"net/http"
)

// Liveness probe returns a 204 when app is live
func (a *Application) Liveness(w http.ResponseWriter, r *http.Request) {
	// Return 204
	w.WriteHeader(http.StatusNoContent)
}

// Readiness probe returns a 204 when app is ready to serve traffic
func (a *Application)  Readiness(w http.ResponseWriter, r *http.Request) {
	// Return 204
	w.WriteHeader(http.StatusNoContent)
}

func (a *Application)  addHealthChecks() {
	a.router.HandleFunc("/liveness", a.Liveness)
	a.router.HandleFunc("/readiness", a.Readiness)
}
