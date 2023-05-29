package loggers

import (
	"context"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const REQUEST_ID_KEY = "request_id"

type CoreLogger struct {
	*logrus.Logger
	Ctx context.Context
}

var coreLogger *CoreLogger

func NewLogger(ctx context.Context) *CoreLogger {
	logger := initLogger()

	coreLogger = &CoreLogger{
		Logger: logger,
		Ctx:    ctx,
	}
	return coreLogger
}

func GetLogger(ctx context.Context) *CoreLogger {
	if coreLogger != nil {
		coreLogger.Ctx = ctx
		return coreLogger
	} else {
		return NewLogger(ctx)
	}
}

func (cl *CoreLogger) FormatLog(event string, detail interface{}, key interface{}) *logrus.Entry {
	var reqID = ""
	if cl.Ctx != nil {
		reqID = cl.Ctx.Value(REQUEST_ID_KEY).(string)
	}
	return cl.Logger.WithFields(logrus.Fields{
		"request_id": reqID,
		"detail":     detail,
		"event":      event,
		"key":        key,
	})

}

func initLogger() *logrus.Logger {
	logger := logrus.New()

	path := "./logs/service.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
		rotatelogs.WithRotationSize(2*1024*1024),
	)
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.FatalLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
