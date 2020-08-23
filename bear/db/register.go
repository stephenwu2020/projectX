package db

import "github.com/pkg/errors"

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

func register() {
	setdbhandler(GET_LOGIN_DATA, getLoginData)
	setdbhandler(CREATE_ROLE, createRole)
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
