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
	//GetSingleWishlistQuery(string) (model.Wishlist, error)
}
