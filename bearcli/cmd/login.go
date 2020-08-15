package cmd

import (
	"bearcli/com_ss_pb_proto"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "send login message to bear",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login() {
	fmt.Println("Simulate login proto")
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal("connect failed", err)
	}
	defer conn.Close()

	var uuid uint32 = 123
	login := com_ss_pb_proto.Cs_10010001{
		Uuid: &uuid,
	}

	data, err := proto.Marshal(&login)
	if err != nil {
		log.Fatal("marshaling failed", err)
	}
	msgId := 10010001

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
	fmt.Println("Sended Reqest:", &login)

	// rece
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("read error:", err)
	}

	recv := &com_ss_pb_proto.Sc_10010001{}
	err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Received response", recv)
}
