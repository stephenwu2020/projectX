package game

import (
	"bear/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Module   = new(GameModule)
)

type GameModule struct {
	*module.Skeleton
}

func (m *GameModule) OnInit() {
	m.Skeleton = skeleton
}

func (m *GameModule) OnDestroy() {

}
