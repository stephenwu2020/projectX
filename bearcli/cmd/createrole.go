package cmd

import (
	"bearcli/com_ss_pb_proto"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

var createRoleCmd = &cobra.Command{
	Use:   "createRole",
	Short: "send create role message to bear",
	Run: func(cmd *cobra.Command, args []string) {
		createRole()
	},
}

func init() {
	rootCmd.AddCommand(createRoleCmd)
}

func createRole() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		log.Fatal("connect failed", err)
	}
	defer conn.Close()

	var uname = "Tom"
	msg := com_ss_pb_proto.Cs_10010002{
		Uname: &uname,
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatal("marshaling failed", err)
	}
	msgId := 10010002

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

	recv := &com_ss_pb_proto.Sc_10010002{}
	err = proto.Unmarshal(buf[lenOfLen+lenOfMsgId:n], recv)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Rece create role response, uid is:", recv.GetUid())
}
