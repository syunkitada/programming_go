package benchmark

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkBuiltinLog(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	log.SetOutput(io.Writer(file))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Print("output log")
	}
}

func getZapConfig() zap.Config {
	// ProductionEncoderConfigからSamplingの設定を外したものと同等
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
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
		OutputPaths: []string{"./output.log"},
	}
}

func BenchmarkZapDevelopmentLog(b *testing.B) {
	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{"./output.log"}
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World")
	}
}

func BenchmarkZapProductionLog(b *testing.B) {
	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"./output.log"}
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World")
	}
}

func BenchmarkZapLog(b *testing.B) {
	conf := getZapConfig()
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World")
	}
}

func BenchmarkZapSamplingLog(b *testing.B) {
	conf := getZapConfig()
	conf.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World")
	}
}

func BenchmarkZapLogDisableCaller(b *testing.B) {
	conf := getZapConfig()
	conf.DisableCaller = true
	conf.DisableStacktrace = true
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World")
	}
}

func BenchmarkZapLogZap(b *testing.B) {
	conf := getZapConfig()
	logger, _ := conf.Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Hello World", zap.String("traceId", "id"), zap.Int("status", 200))
	}
}

func BenchmarkZapSugerfLog(b *testing.B) {
	conf := getZapConfig()
	logger, _ := conf.Build()
	sugarLogger := logger.Sugar()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugarLogger.Infof("Hello World traceId=%s, status=%d", "id", 200)
	}
}

func BenchmarkZapSugerwLog(b *testing.B) {
	conf := getZapConfig()
	logger, _ := conf.Build()
	sugarLogger := logger.Sugar()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugarLogger.Infow("Hello World", "traceId", "id", "status", 200)
	}
}

func BenchmarkZerolog(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logger := zerolog.New(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Msg("Hello World")
	}
}

func BenchmarkZerologUnix(b *testing.B) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logger := zerolog.New(file)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Msg("Hello World")
	}
}
