package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckWishlistExist checks whether Wishlist exist or deleted
func (dc *DBRepo) CheckWishlistExist(productID int) bool {

	collection := dc.DBClient.Database("local").Collection("Wishlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var count int64

	//Getting the count of documents for the supplied ID
	count, err := collection.CountDocuments(ctx, bson.D{primitive.E{Key: "ProductID", Value: productID},
		primitive.E{Key: "IsDeleted", Value: false}})

	log.Println("Total count matched ", count)
	if count <= 0 || err != nil {
		if err != nil {
			log.Println("Check Wishlist exist error ", err)
		}
		return false
	}

	return true
}
