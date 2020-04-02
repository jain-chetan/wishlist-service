package api

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/wishlist-service/lib/config"
)

// //APIHandler structure to group all get methods
// type APIHandler struct{}

//VersionHandler to check wishlist service response
func (api *APIHandler) VersionHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	version := map[string]interface{}{"version": config.Version}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(response).Encode(version)

}
