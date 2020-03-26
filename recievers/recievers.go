package recievers

import (
	api "github.com/jain-chetan/wishlist-service/handlers/api"
	get "github.com/jain-chetan/wishlist-service/handlers/get"
)

//RecieverHandler - All Reciever type Handler struct
type RecieverHandler struct {
	APIHandler *api.APIHandler
	GetHandler *get.GetHandler
}

//Initialization function for RecieverHandler
func Initialization() *RecieverHandler {
	RecieverHandler := new(RecieverHandler)

	RecieverHandler.APIHandler = new(api.APIHandler)
	RecieverHandler.GetHandler = new(get.GetHandler)

	return RecieverHandler
}
