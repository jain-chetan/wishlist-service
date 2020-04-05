package database

import (
	"context"
	"time"

	"github.com/jain-chetan/wishlist-service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateWishlistQuery query function to insert wishlist in database
func (dc *DBRepo) CreateWishlistQuery(wishlist model.Wishlist) (model.CreateResponse, error) {

	var result model.CreateResponse
	emptyResponse := model.CreateResponse{}

	collection := dc.DBClient.Database("local").Collection("Wishlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, wishlist)
	if err != nil {
		return emptyResponse, err
	}

	result.ID = res.InsertedID.(primitive.ObjectID)

	return result, nil
}
