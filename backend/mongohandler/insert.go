package mongohandler

import (
	"context"
	"fmt"
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

	_, readErr := Read(collectionName, object)
	if readErr != nil {
		if readErr != mongo.ErrNoDocuments {
			err := fmt.Errorf("Error checking if object already exists: %w", readErr)
			return nil, err
		}
	} else {
		err := fmt.Errorf("object already exists")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Database.Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, object)

	return insertResult, err
}
