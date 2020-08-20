package game

import (
	"github.com/name5566/leaf/module"
)

type GameModule struct {
	*module.Skeleton
}

func (m *GameModule) OnInit() {
	m.Skeleton = skeleton
}

func (m *GameModule) OnDestroy() {

}
