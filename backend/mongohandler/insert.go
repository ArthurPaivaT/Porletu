package mongohandler

import (
	"context"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(collectionName string, object interface{}) (*mongo.InsertOneResult, error) {

	v := validator.New()
	err := v.Struct(object)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, object)

	return insertResult, err
}
