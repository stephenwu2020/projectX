package db

func register() {
	skeleton.RegisterChanRPC("hello", hello)
	skeleton.RegisterChanRPC("getLoginData", getLoginData)
}
