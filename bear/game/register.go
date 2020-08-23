package game

import (
	"bear/com_ss_pb_proto"
	"reflect"
)

func register() {
	setMsghandler(&com_ss_pb_proto.Cs_10010002{}, handleCreateRole)
}

func setMsghandler(msg interface{}, handler interface{}) {
	Module.RegisterChanRPC(reflect.TypeOf(msg), handler)
}
