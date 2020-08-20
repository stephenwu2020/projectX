package login

import (
	"bear/base"
	"bear/com_ss_pb_proto"
	"reflect"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Module   = new(LoginModule)
)

func init() {
	setMsghandler(&com_ss_pb_proto.Cs_10010001{}, handleLogin)
}

func setMsghandler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
