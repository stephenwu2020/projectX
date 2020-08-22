package e2e

import (
	log "github.com/sirupsen/logrus"
)

var (
	lenOfLen   = 4
	lenOfMsgId = 4
	ip         = "127.0.0.1:3563"
	//ip = "140.82.11.84:3563"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
	})
}
