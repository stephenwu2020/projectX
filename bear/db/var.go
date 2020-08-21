package db

import (
	"bear/base"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	skeleton       = base.NewSkeleton()
	ChanRPC        = skeleton.ChanRPCServer
	Module         = new(DBModule)
	client         *mongo.Client
	DBName         = "bear"
	CollectionUser = "user"
)
