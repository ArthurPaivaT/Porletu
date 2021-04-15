package mongohandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Read(collectionName, objectID string, object interface{}) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	findResult := collection.FindOne(ctx, object)

	return findResult, findResult.Err()
}
