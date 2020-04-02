package gethandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jain-chetan/wishlist-service/interfaces"
	"github.com/jain-chetan/wishlist-service/model"
)

//GetHandler structure to group all get methods
type GetHandler struct{}

//GetSingleWishlistHandler to handle request and response for get single wishlist
func (get *GetHandler) GetSingleWishlistHandler(response http.ResponseWriter, request *http.Request) {

	headerParam := request.Header.Get("userID")

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//converting string ID into int
	userID, err := strconv.Atoi(headerParam)
	if err != nil {
		errResponse := model.Response{
			Code:    400,
			Message: "Invalid Input: unable to convert header from string to int",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	wishlist, err := interfaces.DBClient.GetSingleWishlistQuery(userID)
	if err != nil {
		errResponse := model.Response{
			Code:    400,
			Message: "Error in getting data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}
	response.Header().Add("Status", "200")
	json.NewEncoder(response).Encode(wishlist)
}
