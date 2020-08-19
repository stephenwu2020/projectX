package e2e

import (
	"bear/com_ss_pb_proto"
	"fmt"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestLogin(t *testing.T) {
	fmt.Println("Simulate login proto")
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		t.Fatal("connect failed", err)
	}
	defer conn.Close()

	var uuid uint32 = 123
	login := com_ss_pb_proto.Cs_10010001{
		Uuid: &uuid,
	}

	data, err := proto.Marshal(&login)
	if err != nil {
		t.Fatal("marshaling failed", err)
	}
	msgId := 10010001
	m := encodeMsg(msgId, data)

	_, err = conn.Write(m)
	if err != nil {
		t.Fatal("Write failed", err)
	}
	fmt.Println("Sended Reqest:", &login)

	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal("read error:", err)
	}

	recv := &com_ss_pb_proto.Sc_10010001{}
	err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv)
	if err != nil {
		t.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Received response", recv)
}
