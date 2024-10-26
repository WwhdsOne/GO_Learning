package main

import (
	"github.com/WwhdsOne/Wlog"
	"github.com/WwhdsOne/Wlog/wlcore"
)

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
	ls := &wlcore.LogSummary{
		LogFormatConfig: &wlcore.LogFormatConfig{
			Level:           WLog.DebugLevel,
			Prefix:          "TEST-ZAP-JSON",
			IsJson:          false,
			EncoderLevel:    WLog.CapitalColorLevelEncoder,
			StacktraceLevel: WLog.ErrorLevel,
		},
	}
	newDefaultLogger := WLog.Build(ls)
	WLog.ReplaceDefault(newDefaultLogger)
	loptions := wlcore.Loptions{
		Option: []any{1, "LOL"},
	}
	newDefaultLogger.Info("LOL %d %s", &loptions)
}
