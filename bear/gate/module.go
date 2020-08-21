package gate

import (
	"bear/conf"
	"bear/game"
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
		AgentChanRPC:    game.ChanRPC,
	}
}
