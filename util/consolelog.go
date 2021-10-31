package util

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	uuid "github.com/satori/go.uuid"
	"os"
)

type LogType int

const logFilePerm = 0666
const (
	LogInfo LogType = iota
	LogWarn
	LogErr
	LogData
)

func (lt LogType) String() string {
	return []string{"LogInfo", "LogWarn", "LogErr", "LogData"}[lt]
}

func CreateStdGoKitLog(serviceName string, debug bool) log.Logger {
	f, err := os.OpenFile("service.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, logFilePerm)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(f))
	logger = log.NewSyncLogger(logger)
	logger = log.With(
		logger,
		"service", serviceName,
		"time", log.DefaultTimestampUTC,
		"caller", log.Caller(3),
	)
	if debug {
		logger = level.NewFilter(logger, level.AllowDebug())
	}
	return logger
}

func ConsoleLog(l log.Logger) log.Logger {
	l = log.With(l, "request_id", uuid.NewV4().String())
	return l
}
