package login

import (
	"github.com/name5566/leaf/module"
)

type LoginModule struct {
	*module.Skeleton
}

func (m *LoginModule) OnInit() {
	m.Skeleton = skeleton
	regiser()
}

func (m *LoginModule) OnDestroy() {

}
