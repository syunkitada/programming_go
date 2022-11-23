package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

type Config struct {
	OutputPaths []string
	Level       string
	Encoding    string
}

var conf = Config{
	OutputPaths: []string{"stdout"},
	Level:       "info",
	Encoding:    "json",
}

const (
	LevelDebug = "Debug"
	LevelInfo  = "Info"
	LevelWarn  = "Warn"
)

func NewZapCoreLevel(levelStr string) (level zapcore.Level) {
	switch levelStr {
	case "Debug":
		level = zap.DebugLevel
	case "Info":
		level = zap.InfoLevel
	case "Warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}
	return
}

func Init(conf *Config) {
	zapConf := zap.Config{
		Level:       zap.NewAtomicLevelAt(NewZapCoreLevel(conf.Level)),
		Development: false,
		Encoding:    conf.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: conf.OutputPaths,
	}

	var err error
	logger, err = zapConf.Build()
	if err != nil {
		fmt.Println("Failed to initialize logger", err.Error())
		os.Exit(1)
	}

	sugar = logger.Sugar()
}

func Sync() {
	logger.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Debugf(msg string, fields ...interface{}) {
	sugar.Debugf(msg, fields...)
}

func Infof(msg string, fields ...interface{}) {
	sugar.Infof(msg, fields...)
}

func Warnf(msg string, fields ...interface{}) {
	sugar.Warnf(msg, fields...)
}
func Errorf(msg string, fields ...interface{}) {
	sugar.Errorf(msg, fields...)
}

func Fatalf(msg string, fields ...interface{}) {
	sugar.Fatalf(msg, fields...)
}
