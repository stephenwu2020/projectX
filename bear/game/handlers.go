package game

import (
	"bear/com_ss_pb_proto"
	"bear/db"
	"bear/msg"
	"bear/msg/processor"

	"github.com/name5566/leaf/gate"
	log "github.com/sirupsen/logrus"
)

func handleCreateRole(args []interface{}) {
	message := args[0].(*com_ss_pb_proto.Cs_10010002)
	agent := args[1].(gate.Agent)

	var (
		uid uint32
		ok  bool
	)

	dbres, _ := db.ChanRPC.CallN(db.CREATE_ROLE, message.GetUname(), message.GetSex())
	if err := dbres[1]; err != nil {
		log.Error("create role fail", err)
	} else {
		uid, ok = dbres[0].(uint32)
		if !ok {
			log.Error("assertion uid fail")
		}
	}

	smsg := processor.MsgWithID{
		MsgID: msg.P1001_Create_Role,
		Msg:   &com_ss_pb_proto.Sc_10010002{Uid: &uid},
	}
	agent.WriteMsg(&smsg)
}
