package benchmark

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func BenchmarkBuiltinLog(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	log.SetOutput(io.Writer(file))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Print("Hello World")
	}
}

func BenchmarkBuiltinLogger(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	logger := log.New(io.Writer(file), "", 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Print("Time=" + time.Now().Format(time.RFC3339) + " Hello World")
	}
}

func BenchmarkWrite(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	writer := io.Writer(file)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := buffer.Buffer{}
		buf.AppendString("Hello World\n")
		writer.Write(buf.Bytes())
	}
}

func BenchmarkBufioWrite(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	writer := bufio.NewWriter(file)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer.Write([]byte("Hello World\n"))
	}
	writer.Flush()
}

func BenchmarkBufioWriteBuffer(b *testing.B) {
	file, _ := os.OpenFile("./output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	writer := bufio.NewWriter(file)
	bytes := []byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes = []byte("Hello World\n")
		writer.Write(bytes)
	}
	writer.Flush()
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
	conf.DisableCaller = true
	conf.DisableStacktrace = true
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
