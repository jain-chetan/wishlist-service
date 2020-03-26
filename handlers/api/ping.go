package api

import (
	"encoding/json"
	"net/http"

	"github.com/jain-chetan/wishlist-service/model"
)

//APIHandler structure to group all get methods
type APIHandler struct{}

//PingHandler to check wishlist service response
func (api *APIHandler) PingHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	pingResponse := model.Response{
		Code:    200,
		Message: "Ok",
	}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(response).Encode(pingResponse)

}
