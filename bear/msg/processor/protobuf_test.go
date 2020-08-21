package processor

import (
	"bear/com_ss_pb_proto"
	"testing"

	"github.com/name5566/leaf/chanrpc"
)

func TestNewProcessor(t *testing.T) {
	p := NewProcessor()
	if p.msgInfo == nil {
		t.Errorf("msgInfo not initialize")
	}
	if p.msgID == nil {
		t.Errorf("msgId not initialize")
	}
}

func TestSetByteOrder(t *testing.T) {
	p := NewProcessor()
	p.SetByteOrder(true)
	if p.littleEndian != true {
		t.Errorf("SetByteOrder fail")
	}
}

func TestRegister(t *testing.T) {
	p := NewProcessor()
	_, err := p.Register(10, nil)
	if err == nil {
		t.Errorf("nil proto not detect")
	}
	msgId, err := p.Register(10, &com_ss_pb_proto.Cs_10010001{})
	_, err = p.Register(10, &com_ss_pb_proto.Cs_10010001{})
	if err == nil {
		t.Errorf("duplicate msg register")
	}
	_, ok := p.msgInfo[msgId]
	if !ok {
		t.Errorf("msg not found in map")
	}
}

func TestSetRouter(t *testing.T) {
	p := NewProcessor()
	msg := &com_ss_pb_proto.Cs_10010001{}
	if err := p.SetRouter(msg, &chanrpc.Server{}); err == nil {
		t.Errorf("msg did not register")
	}
	p.Register(10, msg)
	if err := p.SetRouter(msg, &chanrpc.Server{}); err != nil {
		t.Errorf("set router failed")
	}
}

func TestMarshal(t *testing.T) {
	p := NewProcessor()
	if _, err := p.Marshal(nil); err == nil {
		t.Errorf("marshal nil msg")
	}
	if _, err := p.Marshal("hello"); err == nil {
		t.Errorf("marshal not proto msg")
	}
}
