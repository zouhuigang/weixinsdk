package logger

import (
	"path"
	"strings"
	"time"
	zconfig "weixinsdk/src/config"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

/*
等级:
PanicLevel >FatalLevel>ErrorLevel>WarnLevel>InfoLevel>DebugLevel>TraceLevel

// PanicLevel level, highest level of severity. Logs and then calls panic with the
// message passed to Debug, Info, ...
PanicLevel Level = iota
// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
// logging level is set to Panic.
FatalLevel
// ErrorLevel level. Logs. Used for errors that should definitely be noted.
// Commonly used for hooks to send errors to an error tracking service.
ErrorLevel
// WarnLevel level. Non-critical entries that deserve eyes.
WarnLevel
// InfoLevel level. General operational entries about what's going on inside the
// application.
InfoLevel
// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
DebugLevel
// TraceLevel level. Designates finer-grained informational events than the Debug.
TraceLevel
*/

var MyLogger *logrus.Logger

var logLevels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

/*
https://github.com/sirupsen/logrus/blob/08e90462da344fbb3880e8e47a0ddacc37508579/example_basic_test.go
https://github.com/kubernetes-sigs/windows-gmsa/blob/77a78c1440270bc9fbd43911aeb3b87efbedc212/admission-webhook/main.go
加载日志配置:
	logPath:日志路径,
	logFileName:日志文件
	maxAge:文件最大保存时间,
	rotationTime:日志切割时间间隔
	WithMaxAge和WithRotationCount二者只能设置一个，
*/
func Load() error {
	//参数
	logPath := zconfig.CFG.MustValue("logger", "path", "")
	//默认30天
	maxAgeInt := zconfig.CFG.MustInt64("logger", "maxAge", 720)
	maxAge := time.Hour * time.Duration(maxAgeInt)
	//默认一天一个文件
	rotationTimeInt := zconfig.CFG.MustInt64("logger", "rotation", 24)
	rotationTime := time.Hour * time.Duration(rotationTimeInt)
	//日志等级
	rawLogLevel := zconfig.CFG.MustValue("logger", "level", "debug")

	//创建一个日志实体
	MyLogger = logrus.New()
	MyLogger.Formatter = new(logrus.JSONFormatter)
	MyLogger.Formatter = new(logrus.TextFormatter)                     //default
	MyLogger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	MyLogger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	//MyLogger.Level = logrus.TraceLevel

	//官方有一个ParseLevel
	if logLevel, valid := logLevels[strings.ToLower(rawLogLevel)]; valid {
		MyLogger.SetLevel(logLevel)
	} else {
		MyLogger.SetLevel(logrus.DebugLevel)
	}

	logFileName := "log"
	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),
		//rotatelogs.WithRotationCount(365),         // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		//失败之后，会使用默认的std输出日志
		MyLogger.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	MyLogger.AddHook(lfHook)

	return nil

}
