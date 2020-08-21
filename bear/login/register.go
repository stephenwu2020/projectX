package login

import (
	"bear/com_ss_pb_proto"
	"reflect"
)

func setMsghandler(msg interface{}, handler interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(msg), handler)
}

func regiser() {
	skeleton.RegisterChanRPC("NewAgent", newAgent)
	skeleton.RegisterChanRPC("CloseAgent", closeAgent)

	setMsghandler(&com_ss_pb_proto.Cs_10010001{}, handleLogin)
}

func newAgent(args []interface{}) {

}

func closeAgent(args []interface{}) {

}
