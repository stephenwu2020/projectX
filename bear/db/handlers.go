package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"bear/db/collections"
	"bear/db/conf"

	"go.mongodb.org/mongo-driver/bson"
)

func getLoginData(args []interface{}) []interface{} {
	uid, ok := args[0].(uint32)
	if !ok {
		return []interface{}{nil, errors.New("uid assertion fail")}
	}
	collection := Module.client.Database(conf.DBName).Collection(collections.COLLECTION_USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"uid": uid}
	var result collections.Users
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return []interface{}{nil, err}
	}
	return []interface{}{&result, nil}
}

func createRole(args []interface{}) []interface{} {
	username, ok := args[0].(string)
	if !ok {
		return []interface{}{nil, errors.New("username assertion fail")}
	}
	sex, ok := args[1].(uint32)
	if !ok {
		return []interface{}{nil, errors.New("sex assertion fail")}
	}
	collections := Module.client.Database(conf.DBName).Collection(collections.COLLECTION_USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	uid := uuid.New().ID()

	_, err := collections.InsertOne(ctx, bson.M{"uid": uid, "name": username, "sex": sex})
	if err != nil {
		return []interface{}{nil, errors.WithMessage(err, "insert users fail")}
	}
	return []interface{}{uid, nil}
}
