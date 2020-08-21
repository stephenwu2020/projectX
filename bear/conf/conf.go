package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags | log.Lshortfile

	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 4
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000

	// server conf
	LogLevel    = "debug"
	LogPath     = ""
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     = "0.0.0.0:3563"
	MaxConnNum  = 20000
	ConsolePort int
	ProfilePath string

	// database conf
	DBURI  = "mongodb://localhost:27017"
	DBName = "bear"
)
