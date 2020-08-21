package main

import (
	"bear/conf"
	"bear/db"
	"bear/game"
	"bear/gate"
	"bear/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	log "github.com/sirupsen/logrus"
)

func init() {
	lconf.LogLevel = conf.LeafLogLevel
	lconf.LogPath = conf.LeafLogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.ConsolePort
	lconf.ProfilePath = conf.ProfilePath

	log.SetLevel(log.DebugLevel)
}

func main() {
	leaf.Run(
		db.Module,
		game.Module,
		gate.Module,
		login.Module,
	)
}
