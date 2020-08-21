package db

import (
	"bear/base"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	skeleton       = base.NewSkeleton()
	ChanRPC        = skeleton.ChanRPCServer
	Module         = new(DBModule)
	client         *mongo.Client
	DBName         = "bear"
	DBUser         = "bear"
	DBPasswd       = "123456"
	CollectionUser = "user"
	DBURI          = fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s?authMechanism=SCRAM-SHA-1", DBUser, DBPasswd, DBName)
)
