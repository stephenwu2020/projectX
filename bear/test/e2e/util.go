package e2e

import "encoding/binary"

func int2bytes32(i int) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func encodeMsg(msgId int, data []byte) []byte {
	m := make([]byte, lenOfLen+lenOfMsgId+len(data))
	//binary.BigEndian.PutUint16(m, uint16(2+len(data)))
	copy(m[0:], int2bytes32(lenOfMsgId+len(data)))
	copy(m[lenOfMsgId:], int2bytes32(msgId))
	copy(m[lenOfLen+lenOfMsgId:], data)
	return m
}
