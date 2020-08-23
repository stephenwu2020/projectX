package login

import (
	"bear/com_ss_pb_proto"
	"reflect"
)

func regiser() {
	Module.RegisterChanRPC("NewAgent", newAgent)
	Module.RegisterChanRPC("CloseAgent", closeAgent)

	setMsghandler(&com_ss_pb_proto.Cs_10010001{}, handleLogin)
}

func setMsghandler(msg interface{}, handler interface{}) {
	Module.RegisterChanRPC(reflect.TypeOf(msg), handler)
}

// event fire at connection establish
func newAgent(args []interface{}) {

}

func closeAgent(args []interface{}) {

}
