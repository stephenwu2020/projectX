package login

import (
	"bear/com_ss_pb_proto"
	"reflect"

	"github.com/name5566/leaf/module"
)

type LoginModule struct {
	*module.Skeleton
}

func (m *LoginModule) OnInit() {
	m.Skeleton = skeleton
	m.setMsghandler(&com_ss_pb_proto.Cs_10010001{}, handleLogin)
}

func (m *LoginModule) OnDestroy() {

}

func (m *LoginModule) setMsghandler(msg interface{}, handler interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), handler)
}
