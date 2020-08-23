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

// Here is the place to establish relationship between socket agent and user
func newAgent(args []interface{}) {

}

// Here release connection between socket and user
func closeAgent(args []interface{}) {

}
