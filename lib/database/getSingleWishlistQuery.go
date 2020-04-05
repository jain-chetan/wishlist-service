package database

import (
	"context"
	"log"
	"time"

	"github.com/jain-chetan/wishlist-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetSingleWishlistQuery gets Wishlist details based on supplied WishlistID
func (dc *DBRepo) GetSingleWishlistQuery(userID int) ([]model.Wishlist, error) {

	var wishlists []model.Wishlist

	collection := dc.DBClient.Database("local").Collection("Wishlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "UserID", Value: 1}})

	//Call to the database to fetch the data
	cursor, err := collection.Find(ctx, bson.D{primitive.E{Key: "UserID", Value: userID},
		primitive.E{Key: "IsDeleted", Value: false}}, findOptions)

	// error handling
	if err != nil {
		log.Println("Error in get Single Wishlist ", err)
		return wishlists, err
	}

	for cursor.Next(ctx) {
		var wishlist model.Wishlist
		err := cursor.Decode(&wishlist)
		if err != nil {
			log.Println("Error in decoding cursor ", err)
			return wishlists, err
		}
		wishlists = append(wishlists, wishlist)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return wishlists, err
	}

	//Returning the data
	return wishlists, nil

}
