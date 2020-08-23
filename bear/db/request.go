package db

type Request struct {
	Data   []interface{}
	Err    error
	Result interface{}
}
