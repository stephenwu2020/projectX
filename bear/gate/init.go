package gate

import (
	"bear/com_ss_pb_proto"
	"bear/game"
	"bear/msg"
)

var (
	Module = new(GateModule)
)

func init() {
	route()
}

func route() {
	msg.Processor.SetRouter(&com_ss_pb_proto.Cs_10010001{}, game.ChanRPC)
	msg.Processor.SetRouter(&com_ss_pb_proto.Cs_10010002{}, game.ChanRPC)
}
