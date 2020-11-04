package logger

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

const (
	Splunk = 1 << iota
	Stdout
)

type Logger struct {
	logger.LogLevel
	OutType int
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	if os.Getenv("env") == "prod" {
		newlogger.OutType = Splunk
	} else {
		newlogger.OutType = Splunk | Stdout
	}
	return &newlogger
}

func (l *Logger) Info(ctx context.Context, s string, i ...interface{}) {
}

func (l *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
}

func (l *Logger) Error(ctx context.Context, s string, i ...interface{}) {
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rowAffected := fc()
}

func (l *Logger) Print(v ...interface{}) {
	fmt.Println("=--------------------=")
	fmt.Fprint(os.Stdout, v)
}

func (l *Logger) Log(s string, i ...interface{}) {
	if l.OutType&Splunk == Splunk {
		go func() {}()
	}
	if l.OutType&Stdout == Stdout {
		go func() {}()
	}
}

type LogMessage struct {
	Method   string `json:"_method"`
	Message  string `json:"msg"`
	From     string `json:"_from"`
	Ms       int    `json:"_ms"`
	UserId   int    `json:"_userId"`
	Name     string `json:"name"`
	Pid      int    `json:"pid"`
	TraceId  string `json:"traceId"`
	ReqUrl   string `json:"req_url"`
	Severity string `json:"severity"`
	Type     string `json:"type"`
	Host     string `json:"host"`
}
