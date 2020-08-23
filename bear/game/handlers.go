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

	var uid uint32

	dbreq := &db.Request{
		Data: []interface{}{message.GetUname(), message.GetSex()},
	}
	if err := db.Module.Request(db.CREATE_ROLE, dbreq); err != nil {
		log.Errorf("Create role fail: %s", err)
	} else if dbreq.Err != nil {
		log.Errorf("Create role fail: %s", dbreq.Err)
	} else {
		uid = dbreq.Result.(uint32)
	}

	smsg := processor.MsgWithID{
		MsgID: msg.P1001_Create_Role,
		Msg:   &com_ss_pb_proto.Sc_10010002{Uid: &uid},
	}
	agent.WriteMsg(&smsg)
}
