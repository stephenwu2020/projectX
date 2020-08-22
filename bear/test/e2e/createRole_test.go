package e2e

import (
	"bear/com_ss_pb_proto"
	"net"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func createRole(conn net.Conn, name *string, sex *uint32) error {
	// send
	msg := com_ss_pb_proto.Cs_10010002{
		Uname: name,
		Sex:   sex,
	}
	data, err := proto.Marshal(&msg)
	if err != nil {
		return errors.WithMessage(err, "Marshal fail")
	}
	msgId := 10010002
	fullmsg := encodeMsg(msgId, data)
	if _, err := conn.Write(fullmsg); err != nil {
		return errors.WithMessage(err, "Write fail")
	}

	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		return errors.WithMessage(err, "Read fail")
	}
	recv := &com_ss_pb_proto.Sc_10010002{}
	if err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv); err != nil {
		return errors.WithMessage(err, "Unmarshal fail")
	}
	log.Infof("Sended:%+v - Received:%+v", &msg, recv)
	return nil
}

func TestCreateRole(t *testing.T) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		t.Fatal("connect failed", err)
	}
	defer conn.Close()
	var (
		name        = "Tom"
		sex  uint32 = 1
	)
	if err := createRole(conn, &name, &sex); err != nil {
		t.Error("Create role failed", err)
	}
	name = "David"
	sex = 2
	if err := createRole(conn, &name, &sex); err != nil {
		t.Error("Create role failed", err)
	}
}
