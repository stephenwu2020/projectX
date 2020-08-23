package req

import (
	"bear/db/dbdata"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
)

func GetLoginData(dbreq *Request) {
	uid, ok := dbreq.Data[0].(uint32)
	if !ok {
		dbreq.Err = errors.New("Assert uid faid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	filter := bson.M{"uid": uid}
	var result dbdata.User
	if err := dbreq.Collection.FindOne(ctx, filter).Decode(&result); err != nil {
		dbreq.Err = errors.WithMessage(err, "Find user by uid fail")
		return
	}
	dbreq.Result = &result
}

func CreateRole(dbreq *Request) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	uid := uuid.New().ID()
	doc := bson.M{"uid": uid, "name": username, "sex": sex}

	if _, err := dbreq.Collection.InsertOne(ctx, doc); err != nil {
		dbreq.Err = errors.WithMessage(err, "insert users fail")
		return
	}

	dbreq.Result = uid
}
