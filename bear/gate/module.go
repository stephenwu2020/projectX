package gate

import (
	"bear/com_ss_pb_proto"
	"bear/conf"
	"bear/game"
	"bear/login"
	"bear/msg"

	"github.com/name5566/leaf/gate"
)

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.CertFile,
		KeyFile:         conf.KeyFile,
		TCPAddr:         conf.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    login.Module.ChanRPCServer,
	}
	m.route()
}

func (m *GateModule) route() {
	msg.Processor.SetRouter(&com_ss_pb_proto.Cs_10010001{}, login.Module.ChanRPCServer)
	msg.Processor.SetRouter(&com_ss_pb_proto.Cs_10010002{}, game.Module.ChanRPCServer)
}
