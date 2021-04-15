package mongohandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(collectionName string, object interface{}) (*mongo.DeleteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	deleteResult, err := collection.DeleteOne(ctx, object)

	return deleteResult, err
}
