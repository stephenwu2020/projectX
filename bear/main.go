package main

import (
	"bear/conf"
	"bear/db"
	"bear/game"
	"bear/gate"
	"bear/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.LogLevel
	lconf.LogPath = conf.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.ConsolePort
	lconf.ProfilePath = conf.ProfilePath

	leaf.Run(
		db.Module,
		game.Module,
		gate.Module,
		login.Module,
	)
}
