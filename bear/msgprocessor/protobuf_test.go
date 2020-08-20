package msgprocessor

import "testing"

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
