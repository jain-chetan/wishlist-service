package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteWishlistQuery deletes wishlist record
func (dc *DBRepo) DeleteWishlistQuery(productID int) (int, error) {

	collection := dc.DBClient.Database("local").Collection("Wishlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	setDeleteFlag := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "IsDeleted", Value: true},
		}},
	}

	res, errUpdate := collection.UpdateOne(ctx, bson.D{primitive.E{Key: "ProductID", Value: productID}}, setDeleteFlag)

	if errUpdate != nil {
		return 0, errUpdate
	}

	return int(res.ModifiedCount), nil
}
