package post

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jain-chetan/wishlist-service/interfaces"
	"github.com/jain-chetan/wishlist-service/model"
)

//PostHandler structure to group all get methods
type PostHandler struct{}

//PostWishlistHandler to handle JSON request and response for database insertion of Wishlist
func (post *PostHandler) PostWishlistHandler(response http.ResponseWriter, request *http.Request) {

	//wishlist model to decode JSON data
	var wishlist model.Wishlist

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Decoding JSON Body from Request
	errDecode := json.NewDecoder(request.Body).Decode(&wishlist)

	//Error Handling for Decoding JSON Request
	if errDecode != nil {
		log.Println("Error in Decoding JSON Body", errDecode)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Decoding JSON Body",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//Calling CreateWishlistQuery to insert wishlist in database
	result, err := interfaces.DBClient.CreateWishlistQuery(wishlist)
	if err != nil {
		log.Println("Error in inserting data ", err)
		response.WriteHeader(http.StatusBadRequest)
		errResponse := model.Response{
			Code:    400,
			Message: "Error in Inserting Data",
		}
		response.Header().Add("Status", "400")
		json.NewEncoder(response).Encode(errResponse)
		return
	}

	//Returning created ID back for new wishlist inserted.
	log.Println("Inserted ID is ", result.ID)
	result = model.CreateResponse{
		ID:      result.ID,
		Code:    201,
		Message: "Successfully Created",
	}
	response.Header().Add("Status", "201")
	json.NewEncoder(response).Encode(result)
}
