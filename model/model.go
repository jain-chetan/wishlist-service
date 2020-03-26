package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Response defines response code and message
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//CreateResponse for insertion response
type CreateResponse struct {
	ID      primitive.ObjectID `json:"wishlistID"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
}

//Wishlist structure for Wishlist document
type Wishlist struct {
	WishlistID primitive.ObjectID `json:"wishlistID,omitempty" bson:"wishlistID,omitempty"`
	UserID     int                `json:"userID" bson:"userID"`
	ProductID  int                `json:"productID" bson:"productID"`
	IsDeleted  bool               `json:"isDeleted" bson:"isDeleted"`
}

//DBConfig has information required to connect to DB
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}
