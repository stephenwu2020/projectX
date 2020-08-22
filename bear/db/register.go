package db

const (
	GET_LOGIN_DATA = iota
)

func register() {
	skeleton.RegisterChanRPC("getLoginData", getLoginData)
}
