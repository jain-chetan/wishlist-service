package recievers

import (
	api "github.com/jain-chetan/wishlist-service/handlers/api"
	delete "github.com/jain-chetan/wishlist-service/handlers/delete"
	get "github.com/jain-chetan/wishlist-service/handlers/get"
	post "github.com/jain-chetan/wishlist-service/handlers/post"
)

//RecieverHandler - All Reciever type Handler struct
type RecieverHandler struct {
	APIHandler    *api.APIHandler
	GetHandler    *get.GetHandler
	PostHandler   *post.PostHandler
	DeleteHandler *delete.DeleteHandler
}

//Initialization function for RecieverHandler
func Initialization() *RecieverHandler {
	RecieverHandler := new(RecieverHandler)

	RecieverHandler.APIHandler = new(api.APIHandler)
	RecieverHandler.GetHandler = new(get.GetHandler)
	RecieverHandler.PostHandler = new(post.PostHandler)
	RecieverHandler.DeleteHandler = new(delete.DeleteHandler)

	return RecieverHandler
}
