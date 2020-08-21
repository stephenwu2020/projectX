package game

import (
	"bear/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Module   = new(GameModule)
)
