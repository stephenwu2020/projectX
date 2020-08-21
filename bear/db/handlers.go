package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func hello(args []interface{}) interface{} {
	return "hello"
}

func getLoginData(args []interface{}) []interface{} {
	uid := args[0].(string)
	collection := client.Database(DBName).Collection(CollectionUser)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"uid": uid}
	var result struct {
		Uid string
	}
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return []interface{}{nil, err}
	}
	return []interface{}{&result, nil}
}
