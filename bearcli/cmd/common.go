package cmd

import "encoding/binary"

var (
	lenOfLen   = 4
	lenOfMsgId = 4
	//ip         = "127.0.0.1:3563"
	ip = "140.82.11.84:3563"
)

func int2bytes32(i int) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}
