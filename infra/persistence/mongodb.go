package persistence

import (
	"context"
	"crud/infra/conf"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	host, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Env.MongoUri))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connect mongo successfully")
	return host.Database(conf.Env.MongoDbName)

}
