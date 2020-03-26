package gethandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/wishlist-service/model"
)

//GetHandler structure to group all get methods
type GetHandler struct{}

//GetSingleWishlistHandler to handle request and response for get single wishlist
func (get *GetHandler) GetSingleWishlistHandler(response http.ResponseWriter, request *http.Request) {

	pathParam := mux.Vars(request)
	wishlistID := pathParam["wishlistID"]

	log.Println("Path parameter: ", wishlistID)

	//wishlist, err := interfaces.DBClient.GetSingleWishlistQuery(wishlistID)
	var err error
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
	json.NewEncoder(response).Encode(err)
}
