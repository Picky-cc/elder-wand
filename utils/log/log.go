package log

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

type LoggerLevel int

const (
	ErrorLevel LoggerLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

var Development = true
var Level = InfoLevel
var FileName string
var ErrorFileName string
var FileMaxSize = 100 // MB
var MaxAge = 7        // day

var LogFileIOWriter *lumberjack.Logger
var ErrorFileIOWriter *lumberjack.Logger

func Init() {
	if zapLogger != nil {
		return
	}
	config()
}

func config() {
	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-temporary outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to temporary and writing
	// them to the console. We'd like to encode the console output and the temporary
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	var lowPriority zap.LevelEnablerFunc
	lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.ErrorLevel
	})

	// Assume that we have clients for two files. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	var topicDebugging zapcore.WriteSyncer
	var topicErrors zapcore.WriteSyncer
	if FileName == "" {
		panic(errors.New("temporary name is empty"))
	}
	if ErrorFileName == "" {
		panic(errors.New("*errors.Error temporary name is empty"))
	}

	LogFileIOWriter = getFileIOWriter(fileTypeLog)
	ErrorFileIOWriter = getFileIOWriter(fileTypeError)
	topicDebugging = zapcore.AddSync(LogFileIOWriter)
	topicErrors = zapcore.AddSync(ErrorFileIOWriter)

	// High-priority output should also go to standard *errors.Error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the temporary output for machine consumption and the console output
	// for human operators.
	fileEncoder := zapcore.NewJSONEncoder(fileEncodeConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncodeConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(fileEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	if Development {
		zapLogger = zap.New(core, zap.AddStacktrace(zap.WarnLevel))
	} else {
		zapLogger = zap.New(core, zap.AddStacktrace(zap.WarnLevel), zap.Development())
	}
}

func Info(msg string, fields ...map[string]interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Info(msg, fields)
}

func Infow(msg string, keyValues ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Infow(msg, keyValues...)
}

func Infof(msg string, v ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	sugar := zapLogger.Sugar()
	sugar.Info(msg)
}

func Debug(msg string, fields ...map[string]interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Debug(msg, fields)
}

func Debugw(msg string, keyValues ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Debugw(msg, keyValues...)
}

func Debugf(msg string, v ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	sugar := zapLogger.Sugar()
	sugar.Debug(msg)
}

func Warn(msg string, fields ...map[string]interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Warn(msg, fields)
}

func Warnw(msg string, keyValues ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Warnw(msg, keyValues...)
}

func Warnf(msg string, v ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	sugar := zapLogger.Sugar()
	sugar.Warn(msg)
}

func Error(msg string, fields ...map[string]interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Error(msg, fields)
}

func Errorw(msg string, keyValues ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()
	sugar := zapLogger.Sugar()
	sugar.Errorw(msg, keyValues...)
}

func Errorf(msg string, v ...interface{}) {
	if zapLogger == nil {
		panic(errors.Errorf("Must init log first"))
	}
	defer zapLogger.Sync()

	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	sugar := zapLogger.Sugar()
	sugar.Error(msg)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05.000 -0700"))
}

// func convertLogLevel() zapcore.Level {
// 	switch Level {
// 	case InfoLevel:
// 		return zapcore.InfoLevel
// 	case DebugLevel:
// 		return zapcore.DebugLevel
// 	case WarnLevel:
// 		return zapcore.WarnLevel
// 	case ErrorLevel:
// 		return zapcore.ErrorLevel
// 	default:
// 		return zapcore.ErrorLevel
// 	}
// }

func fileEncodeConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func consoleEncodeConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

type fileLogType int

const (
	fileTypeLog fileLogType = iota
	fileTypeError
)

func getFileIOWriter(logType fileLogType) *lumberjack.Logger {
	var fileName string
	switch logType {
	case fileTypeLog:
		fileName = FileName
	case fileTypeError:
		fileName = ErrorFileName
	default:
		fileName = FileName
	}
	var ioWriter = &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    FileMaxSize, // MB
		MaxBackups: 10,          // number of backups
		MaxAge:     MaxAge,      //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}
	return ioWriter
}
