package db

import (
	"context"
	"errors"
	"time"

	"bear/db/collections"
	"bear/db/conf"

	"go.mongodb.org/mongo-driver/bson"
)

func getLoginData(args []interface{}) []interface{} {
	uid, ok := args[0].(uint32)
	if !ok {
		return []interface{}{nil, errors.New("uid assetion fail")}
	}
	collection := client.Database(conf.DBName).Collection(collections.COLLECTION_USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"uid": uid}
	var result collections.Users
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return []interface{}{nil, err}
	}
	return []interface{}{&result, nil}
}
