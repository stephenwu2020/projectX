package login

import (
	"bear/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Module   = new(LoginModule)
)
