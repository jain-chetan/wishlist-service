package interfaces

import (
	"github.com/jain-chetan/wishlist-service/model"
)

//DBClient object
var DBClient DBInteractions

//-----------------------------------------------------------------------------------------------------------------------//

//DBInteractions interface contains database methods
type DBInteractions interface {
	DBConnect(model.DBConfig) error
	GetSingleWishlistQuery(int) ([]model.Wishlist, error)
	CreateWishlistQuery(model.Wishlist) (model.CreateResponse, error)
	CheckWishlistExist(int) bool
	DeleteWishlistQuery(int) (int, error)
}
