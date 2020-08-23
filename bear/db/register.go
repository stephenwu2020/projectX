package db

import "github.com/pkg/errors"

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

func register() {
	//Module.RegisterChanRPC(GET_LOGIN_DATA, getLoginData)
	setdbhandler(GET_LOGIN_DATA, getLoginData)
	Module.RegisterChanRPC(CREATE_ROLE, createRole)
}

func setdbhandler(msgtype string, handler func(*Request)) {
	fn := func(args []interface{}) {
		dbreq, ok := args[0].(*Request)
		if !ok {
			panic(errors.New("Assert db request fail"))
		}
		handler(dbreq)
	}
	Module.RegisterChanRPC(msgtype, fn)
}
