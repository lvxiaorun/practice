package main

import (
	"github.com/lvxiaorun/logger"
	"go.uber.org/zap"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":6745")
	if err != nil {
		logger.Error("listen", zap.Error(err))
		return
	}
	defer listen.Close()
	con, err := listen.Accept()
	if err != nil {
		logger.Error("accept", zap.Error(err))
		return
	}
	defer con.Close()
	buf := make([]byte, 1024)
	con.Read(buf)
	logger.Debug("res", zap.String("buf", string(buf)))
	con.Write([]byte("HTTP/1.1 400 OK" + "\n\n" + "hahaha"))
}
