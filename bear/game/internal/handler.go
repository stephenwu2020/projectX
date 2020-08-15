package internal

import (
	"bear/com_ss_pb_proto"
	"bear/encode"
	"bear/msg"
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&com_ss_pb_proto.Cs_10010001{}, handleLogin)
	handler(&com_ss_pb_proto.Cs_10010002{}, handleCreateRole)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	m := args[0].(*com_ss_pb_proto.Cs_10010001)
	a := args[1].(gate.Agent)
	log.Debug("Rece login request, uuid is: %v", m.GetUuid())
	res := true

	smsg := encode.ServerMsg{
		MsgID: msg.P1001_LOGIN,
		Msg:   &com_ss_pb_proto.Sc_10010001{LoginResult: &res},
	}
	a.WriteMsg(&smsg)
}

func handleCreateRole(args []interface{}) {
	m := args[0].(*com_ss_pb_proto.Cs_10010002)
	a := args[1].(gate.Agent)
	log.Debug("Rece create role request, uname is: %v", m.GetUname())

	var uid uint32 = 12345678
	smsg := encode.ServerMsg{
		MsgID: msg.P1001_Create_Role,
		Msg:   &com_ss_pb_proto.Sc_10010002{Uid: &uid},
	}
	a.WriteMsg(&smsg)
}
