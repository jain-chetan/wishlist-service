package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/wishlist-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetSingleWishlistQuery gets Wishlist details based on supplied WishlistID
func (dc *DBRepo) GetSingleWishlistQuery(userID int) (model.Wishlist, error) {

	collection := dc.DBClient.Database("local").Collection("Wishlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wishlist model.Wishlist

	//Call to the database to fetch the data
	err := collection.FindOne(ctx, bson.D{primitive.E{Key: "UserID", Value: userID},
		primitive.E{Key: "IsDeleted", Value: false}}).Decode(&wishlist)

	// error handling
	if err != nil {
		log.Println("Error in get Single Wishlist ", err)
		return wishlist, err
	}

	//Returning the data
	return wishlist, nil

}
