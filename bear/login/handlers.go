package login

import (
	"bear/com_ss_pb_proto"
	"bear/db"
	"bear/db/collections"
	"bear/msg"
	"bear/msg/processor"

	"github.com/name5566/leaf/gate"
	log "github.com/sirupsen/logrus"
)

func handleLogin(args []interface{}) {
	message := args[0].(*com_ss_pb_proto.Cs_10010001)
	agent := args[1].(gate.Agent)
	uid := message.GetUid()

	res := false

	dbres, _ := db.ChanRPC.CallN(db.GET_LOGIN_DATA, uid)

	if err := dbres[1]; err != nil {
		log.Error("call db method failed: %s", err)
	} else {
		users, ok := dbres[0].(collections.Users)
		if !ok {
			log.Error("assertion users fail")
		} else if users.Uid == uid {
			res = true
		}
	}

	smsg := processor.MsgWithID{
		MsgID: msg.P1001_LOGIN,
		Msg:   &com_ss_pb_proto.Sc_10010001{LoginResult: &res},
	}
	agent.WriteMsg(&smsg)
}
