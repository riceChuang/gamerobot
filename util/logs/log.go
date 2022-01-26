package logs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	mlog    *logrus.Logger
	logConf *LogConf
	logOnce sync.Once
)

func GetLogger() *logrus.Logger {
	if mlog == nil {
		InitializeLogger(LogConf{
			ENV:       "production",
			Path:      defaultLogDirName,
			LogLevel:  "info",
			LogPretty: false,
		})
	}
	return mlog
}

type LogConf struct {
	ENV       string
	Path      string
	LogLevel  string
	LogPretty bool
}

///初始化日志
func InitializeLogger(conf LogConf) {
	logOnce.Do(func() {
		logConf = &conf
		if logConf.Path == "" {
			logConf.Path = defaultLogDirName
		}

		logLevel, err := logrus.ParseLevel(conf.LogLevel)
		if err != nil {
			logLevel = logrus.DebugLevel
		}

		//初始化日志对象
		mlog = logrus.New()
		mlog.SetLevel(logLevel)
		mlog.SetReportCaller(true)

		mlog.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: logConf.LogPretty,
		})

		mlog.SetFormatter(&myFormatter{
			logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "2006-01-02 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: true,
			},
		})

		//创建Hook
		for level := logrus.PanicLevel; level <= logLevel; level++ {
			hook, err := NewLogrusFileLoggerHook(logConf.Path, maxLogFileSize, level)
			if nil != err {
				panic(err)
			}
			mlog.AddHook(hook)
		}

		//开启日志清理器
		logFileCleaner = NewLogFileCleaner(logConf.Path, maxLogFileSize, maxLogDirSize)
		go logFileCleaner.CheckLogFileSizeAndRemove()
	})
}

func ReloadLevel(logLevel string) {
	if mlog == nil {
		return
	}

	if logConf.LogLevel == logLevel {
		return
	}

	mlog.WithFields(logrus.Fields{
		"conf":            logConf,
		"input_log_level": logLevel,
	}).Info("reload Level Before")

	ll, err := logrus.ParseLevel(logLevel)
	if err != nil {
		mlog.WithFields(logrus.Fields{
			"conf":            logConf,
			"input_log_level": logLevel,
		}).Error(err)
		ll = logrus.DebugLevel
	} else {
		logConf.LogLevel = logLevel
	}

	mlog.WithFields(logrus.Fields{
		"conf":            logConf,
		"input_log_level": logLevel,
	}).Info("reload Level After")

	mlog.SetLevel(ll)
}

type myFormatter struct {
	logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 32 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}


	formatMsg := fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - func:%s - file:%s:%d - %s\n",
		entry.Time.Format(f.TimestampFormat),
		levelColor,
		strings.ToUpper(entry.Level.String()),
		entry.Caller.Function,
		entry.Caller.File,
		entry.Caller.Line,
		entry.Message)

	return []byte(formatMsg), nil
}
