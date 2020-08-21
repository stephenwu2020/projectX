package game

import (
	"bear/com_ss_pb_proto"
	"reflect"
)

func register() {
	skeleton.RegisterChanRPC("NewAgent", newAgent)
	skeleton.RegisterChanRPC("CloseAgent", closeAgent)

	setMsghandler(&com_ss_pb_proto.Cs_10010002{}, handleCreateRole)
}

func setMsghandler(msg interface{}, handler interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), handler)
}

func newAgent(args []interface{}) {

}

func closeAgent(args []interface{}) {

}
