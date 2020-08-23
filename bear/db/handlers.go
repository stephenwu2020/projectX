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

func getLoginData(dbreq *Request) {
	uid, ok := dbreq.Data[0].(uint32)
	if !ok {
		dbreq.Err = errors.New("Assert uid faid")
		return
	}

	collection := Module.client.Database(conf.DBName).Collection(collections.COLLECTION_USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"uid": uid}
	var result collections.Users
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		dbreq.Err = errors.WithMessage(err, "Find user by uid fail")
		return
	}
	dbreq.Result = &result
}

func createRole(dbreq *Request) {
	username, ok := dbreq.Data[0].(string)
	if !ok {
		dbreq.Err = errors.New("Assert username faid")
		return
	}
	sex, ok := dbreq.Data[1].(uint32)
	if !ok {
		dbreq.Err = errors.New("Assert sex faid")
		return
	}
	collections := Module.client.Database(conf.DBName).Collection(collections.COLLECTION_USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	uid := uuid.New().ID()

	if _, err := collections.InsertOne(ctx, bson.M{"uid": uid, "name": username, "sex": sex}); err != nil {
		dbreq.Err = errors.WithMessage(err, "insert users fail")
	}

	dbreq.Result = uid
}
