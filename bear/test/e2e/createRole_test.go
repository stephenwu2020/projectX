package e2e

import (
	"bear/com_ss_pb_proto"
	"fmt"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCreateRole(t *testing.T) {
	fmt.Println("Simulate role proto")

	conn, err := net.Dial("tcp", ip)
	if err != nil {
		t.Fatal("connect failed", err)
	}
	defer conn.Close()

	var uname = "Tom"
	msg := com_ss_pb_proto.Cs_10010002{
		Uname: &uname,
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		t.Fatal("marshaling failed", err)
	}
	msgId := 10010002

	// m: len + msgId + proto
	m := encodeMsg(msgId, data)

	_, err = conn.Write(m)
	if err != nil {
		t.Fatal("Write failed", err)
	}

	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal("read error:", err)
	}
	fmt.Println("Sended Reqest:", &msg)

	recv := &com_ss_pb_proto.Sc_10010002{}
	err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv)
	if err != nil {
		t.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Received response", recv)
}
