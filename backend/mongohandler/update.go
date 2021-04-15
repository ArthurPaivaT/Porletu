package mongohandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Update(collectionName, objectID string, object interface{}) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	updateResult, err := collection.UpdateByID(ctx, objectID, object)

	return updateResult, err
}
