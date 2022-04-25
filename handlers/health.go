package handlers

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// GetServiceHealth returns the availability status of the service
func GetServiceHealth(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Ok Healthy"))
}

// GetServiceVersion returns the version number set in the envionment variable 'SERVICE_VERSION'
func GetServiceVersion(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
    version := os.Getenv("SERVICE_VERSION")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(version))
}
