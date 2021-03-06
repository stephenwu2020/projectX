package db

import (
	"bear/base"
	"bear/db/conf"
	"bear/db/req"
	"context"
	"time"

	"github.com/name5566/leaf/module"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBModule struct {
	*module.Skeleton
	client *mongo.Client
}

func NewDBModule() *DBModule {
	return &DBModule{
		Skeleton: base.NewSkeleton(),
	}
}

/**
***	Lifecycle methods
**/

func (m *DBModule) OnInit() {
	if err := m.connect(); err != nil {
		log.Fatal("Database connect fail: %s", err)
	}
	if err := m.Ping(); err != nil {
		log.Fatal("Database ping fail: %s", err)
	}
	log.Info("Database connect success!")
	m.register()
}

func (m *DBModule) OnDestroy() {
	if err := m.disconnect(); err != nil {
		panic(err)
	}
	log.Info("Database disconnected")
}

/**
*** Public Methods
**/

func (m *DBModule) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return m.client.Ping(ctx, readpref.Primary())
}

func (m *DBModule) Request(msgtype string, dbreq *req.Request) error {
	return m.ChanRPCServer.Call0(msgtype, dbreq)
}

func (m *DBModule) DB() *mongo.Database {
	return m.client.Database(conf.DBName)
}

func (m *DBModule) Collection(name string, opt ...*options.CollectionOptions) *mongo.Collection {
	return m.client.Database(conf.DBName).Collection(name, opt...)
}

/**
*** Private Methods
**/

func (m *DBModule) connect() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DBURI))
	if err != nil {
		return errors.WithMessage(err, "Connect mongodb failed")
	}
	m.client = client

	return nil
}

func (m *DBModule) disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.client.Disconnect(ctx); err != nil {
		return errors.WithMessage(err, "Disconnect mongodb failed")
	}
	return nil
}

func (m *DBModule) register() {
	for _, reqhandler := range requests {
		m.setreq(reqhandler.MsgType, reqhandler.Handler)
	}
}

func (m *DBModule) setreq(msgtype string, handler func(*req.Request)) {
	fn := func(args []interface{}) {
		dbreq, ok := args[0].(*req.Request)
		if !ok {
			panic(errors.New("Assert db request fail"))
		}
		if dbreq.Collection == nil {
			panic(errors.New("Collection not set"))
		}
		handler(dbreq)
	}
	m.RegisterChanRPC(msgtype, fn)
}
