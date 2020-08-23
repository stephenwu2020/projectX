package game

import (
	"bear/com_ss_pb_proto"
	"reflect"
)

func register() {
	Module.RegisterChanRPC("NewAgent", newAgent)
	Module.RegisterChanRPC("CloseAgent", closeAgent)

	setMsghandler(&com_ss_pb_proto.Cs_10010002{}, handleCreateRole)
}

func setMsghandler(msg interface{}, handler interface{}) {
	Module.RegisterChanRPC(reflect.TypeOf(msg), handler)
}

// Must login first, so new agent not allow anywhere
func newAgent(args []interface{}) {

}

func closeAgent(args []interface{}) {

}
