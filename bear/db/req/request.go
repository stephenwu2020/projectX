package req

import "go.mongodb.org/mongo-driver/mongo"

type Request struct {
	Collection *mongo.Collection
	Data       []interface{}
	Err        error
	Result     interface{}
}
