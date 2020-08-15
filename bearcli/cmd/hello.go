package cmd

import (
	"bearcli/msg"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

var (
	lenOfLen   = 4
	lenOfMsgId = 4
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "send hello message to bear",
	Run: func(cmd *cobra.Command, args []string) {
		sayHello()
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

func sayHello() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		log.Fatal("connect failed", err)
	}
	defer conn.Close()

	hello := msg.Hello{
		Name: proto.String("Tom Proto"),
	}
	data, err := proto.Marshal(&hello)
	if err != nil {
		log.Fatal("marshaling failed", err)
	}
	msgId := 1

	// m: len + msgId + proto
	m := make([]byte, lenOfLen+lenOfMsgId+len(data))
	//binary.BigEndian.PutUint16(m, uint16(2+len(data)))
	copy(m[0:], int2bytes32(lenOfMsgId+len(data)))
	copy(m[lenOfMsgId:], int2bytes32(msgId))
	copy(m[lenOfLen+lenOfMsgId:], data)

	_, err = conn.Write(m)
	if err != nil {
		log.Fatal("Write failed", err)
	}

	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("read error:", err)
	}

	recv := &msg.Hello{}
	err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("rece:", recv.GetName())
}

func int2bytes16(i int) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func int2bytes32(i int) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}
