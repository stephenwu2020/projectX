package db

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

func register() {
	Module.RegisterChanRPC(GET_LOGIN_DATA, getLoginData)
	Module.RegisterChanRPC(CREATE_ROLE, createRole)
}
