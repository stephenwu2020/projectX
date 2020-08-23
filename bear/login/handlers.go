package login

import (
	"bear/com_ss_pb_proto"
	"bear/db"
	"bear/db/dbdata"
	"bear/db/req"
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
	smsg := processor.MsgWithID{
		MsgID: msg.P1001_LOGIN,
		Msg:   &com_ss_pb_proto.Sc_10010001{LoginResult: &res},
	}

	dbreq := req.Request{
		Collection: db.Module.Collection(dbdata.COLLECTION_USERS),
		Data:       []interface{}{uid},
	}
	if err := db.Module.Request(db.GET_LOGIN_DATA, &dbreq); err != nil {
		log.Error("Login fail: %s", err)
	} else if dbreq.Err != nil {
		log.Errorf("Login fail: %s", dbreq.Err)
	} else {
		res = true
	}
	agent.WriteMsg(&smsg)
}
