package req

import "go.mongodb.org/mongo-driver/mongo"

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

type Request struct {
	Collection *mongo.Collection
	Data       []interface{}
	Err        error
	Result     interface{}
}

type ModuleReq struct {
	MsgType string
	Handler func(*Request)
}
