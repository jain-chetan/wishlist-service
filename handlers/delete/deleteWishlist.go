package delete

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/wishlist-service/interfaces"
	"github.com/jain-chetan/wishlist-service/model"
)

// DeleteHandler handler to handle delete request
type DeleteHandler struct{}

// DeleteWishlistHandler process delete request and calls delete query
func (deleteData *DeleteHandler) DeleteWishlistHandler(response http.ResponseWriter, request *http.Request) {

	param := mux.Vars(request)
	pathParam := param["productID"]

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//converting string ID into int
	productID, err := strconv.Atoi(pathParam)
	if err != nil {
		errResponse := model.Response{
			Code:    400,
			Message: "Invalid Input: unable to convert header from string to int",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	iswishlistExist := interfaces.DBClient.CheckWishlistExist(productID)

	if !iswishlistExist {
		errResponse := model.Response{
			Code:    400,
			Message: "No Records Found",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	deleteResponse, err := interfaces.DBClient.DeleteWishlistQuery(productID)
	if err != nil || deleteResponse <= 0 {
		log.Println("Error in deleting Data in database ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Deleting data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	result := model.Response{
		Code:    200,
		Message: "Ok",
	}
	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(result)
}
