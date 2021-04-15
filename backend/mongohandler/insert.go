package mongohandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(collectionName string, object interface{}) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, object)

	return insertResult, err
}
