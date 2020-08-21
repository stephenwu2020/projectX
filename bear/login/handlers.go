package login

import (
	"bear/com_ss_pb_proto"
	"bear/db"
	"bear/msg"
	"bear/msg/processor"

	"github.com/name5566/leaf/gate"
	log "github.com/sirupsen/logrus"
)

func handleLogin(args []interface{}) {
	m := args[0].(*com_ss_pb_proto.Cs_10010001)
	a := args[1].(gate.Agent)
	log.Debug("Rece login request, uuid is: %v", m.GetUuid())
	res := true

	hellores, _ := db.ChanRPC.Call1("hello")
	log.Debug("call hello res: %s", hellores)

	dbres, _ := db.ChanRPC.CallN("getLoginData", "123")
	if err := dbres[1].(error); err != nil {
		log.Error("call db method failed: %s", err)
	} else {
		loginres := dbres[0]
		log.Debug("call db res: %v", loginres)
	}

	smsg := processor.MsgWithID{
		MsgID: msg.P1001_LOGIN,
		Msg:   &com_ss_pb_proto.Sc_10010001{LoginResult: &res},
	}
	a.WriteMsg(&smsg)
}
