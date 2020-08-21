package processor

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
)

type Processor struct {
	littleEndian bool
	msgInfo      map[uint32]*MsgInfo
	msgID        map[reflect.Type]uint32
}

type MsgInfo struct {
	msgType   reflect.Type
	msgRouter *chanrpc.Server
}

type MsgRaw struct {
	msgID      uint32
	msgRawData []byte
}

type MsgWithID struct {
	MsgID uint32
	Msg   proto.Message
}

func NewProcessor() *Processor {
	p := new(Processor)
	p.littleEndian = false
	p.msgID = make(map[reflect.Type]uint32)
	p.msgInfo = make(map[uint32]*MsgInfo)
	return p
}

func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (p *Processor) Register(id uint32, msg proto.Message) (uint32, error) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return 0, errors.New("protobuf message pointer required")
	}
	if _, ok := p.msgID[msgType]; ok {
		return 0, fmt.Errorf("message %s is already registered\n", msgType)
	}

	msgInfo := &MsgInfo{
		msgType: msgType,
	}

	p.msgInfo[id] = msgInfo
	p.msgID[msgType] = id
	return id, nil
}

func (p *Processor) SetRouter(msg proto.Message, msgRouter *chanrpc.Server) error {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		return fmt.Errorf("message %s not registered", msgType)
	}

	p.msgInfo[id].msgRouter = msgRouter
	return nil
}

func (p *Processor) Route(msg interface{}, userData interface{}) error {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		return fmt.Errorf("message %s not registered", msgType)
	}
	msgInfo := p.msgInfo[id]
	if msgInfo.msgRouter != nil {
		msgInfo.msgRouter.Go(msgType, msg, userData)
	}
	return nil
}

func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	var id uint32
	if p.littleEndian {
		id = binary.LittleEndian.Uint32(data)
	} else {
		id = binary.BigEndian.Uint32(data)
	}

	msgInfo, ok := p.msgInfo[id]
	if !ok {
		return nil, fmt.Errorf("message id %v not registered", id)
	} else {
		msg := reflect.New(msgInfo.msgType.Elem()).Interface()
		return msg, proto.UnmarshalMerge(data[4:], msg.(proto.Message))
	}
}

func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	smsg, ok := msg.(*MsgWithID)
	if !ok {
		return nil, errors.New("nil message found")
	}
	id := make([]byte, 4)
	if p.littleEndian {
		binary.LittleEndian.PutUint32(id, smsg.MsgID)
	} else {
		binary.BigEndian.PutUint32(id, smsg.MsgID)
	}
	data, err := proto.Marshal(smsg.Msg)
	return [][]byte{id, data}, err
}
