package e2e

import (
	"bear/com_ss_pb_proto"
	"net"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func TestLogin(t *testing.T) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		t.Fatal("connect failed", err)
	}
	defer conn.Close()
	uids := []uint32{4005096660, 488719350, 123}
	for _, uid := range uids {
		t.Run("Login", func(t *testing.T) {
			if err := login(conn, &uid); err != nil {
				t.Error("Login fail", err)
			}
		})
	}
}

func login(conn net.Conn, uid *uint32) error {
	login := &com_ss_pb_proto.Cs_10010001{
		Uid: uid,
	}
	data, err := proto.Marshal(login)
	if err != nil {
		return errors.WithMessage(err, "Marshaling failed")
	}
	msgId := 10010001
	m := encodeMsg(msgId, data)
	_, err = conn.Write(m)
	if err != nil {
		return errors.WithMessage(err, "Write failed")
	}
	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		return errors.WithMessage(err, "Read error:")
	}

	recv := &com_ss_pb_proto.Sc_10010001{}
	if err := proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv); err != nil {
		return errors.WithMessage(err, "Unmarshal fail")
	}

	log.Infof("Sended:%+v - Received:%+v", login, recv)
	return nil

}
