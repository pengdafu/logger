package logger

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"os"
	"time"
)

const (
	Splunk = 1 << iota
	Stdout
)

type Logger struct {
	logger.LogLevel
	OutType       int
	SlowThreshold time.Duration
	ServiceName   string
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
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		sql, rows := fc()
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			if rows == -1 {
				l.Log(
					ctx,
					fmt.Sprintf("Executed: /* %v */ %v ;Elapsed time: %v; error: %v", l.ServiceName,
						sql, float64(elapsed.Nanoseconds())/1e6, err),
				)
			} else {
				l.Log(
					ctx,
					fmt.Sprintf("Executed: /* %v */ %v ;Elapsed time: %v; error: %v", l.ServiceName,
						sql, float64(elapsed.Nanoseconds())/1e6, err),
				)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= logger.Info:
			if rows == -1 {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (l *Logger) Print(v ...interface{}) {
	fmt.Println("=--------------------=")
	fmt.Fprint(os.Stdout, v)
}

func (l *Logger) Log(c context.Context, msg string) {

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
