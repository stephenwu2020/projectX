package game

import (
	"bear/com_ss_pb_proto"
	"reflect"

	"github.com/name5566/leaf/module"
)

type GameModule struct {
	*module.Skeleton
}

func (m *GameModule) OnInit() {
	m.Skeleton = skeleton
	m.setMsghandler(&com_ss_pb_proto.Cs_10010002{}, handleCreateRole)
}

func (m *GameModule) OnDestroy() {

}

func (m *GameModule) setMsghandler(msg interface{}, handler interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), handler)
}
