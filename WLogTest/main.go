package main

import "github.com/WwhdsOne/Wlog/WLog"

/*
localFileWriter:

	fileName: "app.log"
	fileDirPath: "./logs"

kafkaWriter:

	topic: "test-topic"
	partition: 0
	host: "localhost"
	port: 9092
*/
func main() {

	l := WLog.Default()
	lo := &WLog.Loptions{
		Package: "testPackage",
		Option:  []any{"LOL", 123},
	}
	l.Debug("Debug %s %d", lo)
	l.Info("Info %s %d", lo)
	l.Warn("Warn %s %d", lo)
}
