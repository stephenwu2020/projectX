package game

import (
	"bear/base"

	"github.com/name5566/leaf/module"
)

type GameModule struct {
	*module.Skeleton
}

func NewGameModule() *GameModule {
	return &GameModule{
		Skeleton: base.NewSkeleton(),
	}
}

func (m *GameModule) OnInit() {
	register()
}

func (m *GameModule) OnDestroy() {

}
