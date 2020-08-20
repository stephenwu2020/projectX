package login

import (
	"bear/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Module   = new(LoginModule)
)

type LoginModule struct {
	*module.Skeleton
}

func (m *LoginModule) OnInit() {
	m.Skeleton = skeleton
}

func (m *LoginModule) OnDestroy() {

}
