package mongohandler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func Connect(resCh chan error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUsr := os.Getenv("MONGOUSR")
	mongoPwd := os.Getenv("MONGOPWD")
	connURL := fmt.Sprintf("mongodb+srv://%s:%s@porletu-test.3mler.mongodb.net/?retryWrites=true&w=majority", mongoUsr, mongoPwd)

	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connURL))
	if err != nil {
		resCh <- err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("Connected to Mongo")

	Database = MongoClient.Database("Main")
}
