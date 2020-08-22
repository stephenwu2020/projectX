package db

const (
	GET_LOGIN_DATA = "getLoginData"
	CREATE_ROLE    = "createRole"
)

func register() {
	skeleton.RegisterChanRPC(GET_LOGIN_DATA, getLoginData)
	skeleton.RegisterChanRPC(CREATE_ROLE, createRole)
}
