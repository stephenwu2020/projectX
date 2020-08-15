package msg

import (
	"bear/com_ss_pb_proto"
	"bear/encode"
)

var (
	Processor                = encode.NewProcessor()
	P1001_LOGIN       uint32 = 10010001
	P1001_Create_Role uint32 = 10010002
)

func init() {
	Processor.Register(P1001_LOGIN, &com_ss_pb_proto.Cs_10010001{})
	Processor.Register(P1001_Create_Role, &com_ss_pb_proto.Cs_10010002{})
}
