package db

import (
	"context"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDBConnection(context context.Context) (*mongo.Client, string) {

	clientOptions, dbName, err := config.GetMongoURI()

	if err != nil {
		fmt.Println("No client Options were provided for mongodb " + err.Error())
		return nil, ""
	}

	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		fmt.Println("Unable to connect to mongo db" + err.Error())
		return nil, ""
	}

	err = client.Ping(context, nil)
	if err != nil {
		fmt.Println("An invalid connection was created to mongodb" + err.Error())
		return nil, ""
	}
	return client, dbName
}
