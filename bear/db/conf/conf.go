package conf

import "fmt"

var (
	DBName   = "bear"
	DBUser   = "bear"
	DBPasswd = "123456"
	DBURI    = fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s?authMechanism=SCRAM-SHA-1", DBUser, DBPasswd, DBName)
)
