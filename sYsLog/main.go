package main

import (
	"log"
	"log/syslog"
)

func main() {
	// 创建一个 Syslog 写入器
	writer, err := syslog.Dial("udp", "47.93.83.136:515", syslog.LOG_INFO|syslog.LOG_LOCAL0, "myapp")
	if err != nil {
		log.Fatalf("无法连接到 Syslog 服务器: %s", err)
	}
	defer writer.Close()

	// 记录日志消息
	err = writer.Info("这是一条 INFO 级别的日志消息")
	if err != nil {
		log.Fatalf("无法写入日志消息: %s", err)
	}

	err = writer.Warning("这是一条 WARNING 级别的日志消息")
	if err != nil {
		log.Fatalf("无法写入日志消息: %s", err)
	}

	err = writer.Err("这是一条 ERROR 级别的日志消息")
	if err != nil {
		log.Fatalf("无法写入日志消息: %s", err)
	}
}
