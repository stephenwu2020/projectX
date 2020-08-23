package login

import (
	"bear/base"

	"github.com/name5566/leaf/module"
)

type LoginModule struct {
	*module.Skeleton
}

func NewLoginModule() *LoginModule {
	return &LoginModule{
		Skeleton: base.NewSkeleton(),
	}
}

func (m *LoginModule) OnInit() {
	regiser()
}

func (m *LoginModule) OnDestroy() {

}
