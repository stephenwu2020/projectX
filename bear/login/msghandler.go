package login

import (
	"bear/com_ss_pb_proto"
	"bear/msg"
	"bear/msg/processor"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleLogin(args []interface{}) {
	m := args[0].(*com_ss_pb_proto.Cs_10010001)
	a := args[1].(gate.Agent)
	log.Debug("Rece login request, uuid is: %v", m.GetUuid())
	res := true

	smsg := processor.MsgWithID{
		MsgID: msg.P1001_LOGIN,
		Msg:   &com_ss_pb_proto.Sc_10010001{LoginResult: &res},
	}
	a.WriteMsg(&smsg)
}
