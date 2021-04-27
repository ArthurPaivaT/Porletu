package mongohandler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//Read tryes to find one document at collectionName, returns it at SigleResult. If no document matches the filter, err is ErrNoDocuments
func Read(collectionName string, filter interface{}) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	findResult := collection.FindOne(ctx, filter)

	return findResult, findResult.Err()
}
