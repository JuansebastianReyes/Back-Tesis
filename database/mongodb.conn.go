package database

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

func GetCollection(collection string) *mongo.Collection {
	Init()
	var (
		usr  = os.Getenv("MONGO_USER")
		pwd  = os.Getenv("MONGO_PASSWORD")
		host = os.Getenv("MONGO_HOST")
		db   = os.Getenv("MONGO_DB")
	)

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", usr, pwd, host, db)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(db).Collection(collection)
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}
