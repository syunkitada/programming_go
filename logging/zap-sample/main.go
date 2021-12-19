package main

import (
	"time"

	"github.com/rs/xid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// https://pkg.go.dev/go.uber.org/zap#Config
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // 有効なログレベル
		Development: false,                               // スタックトレースの粒度が変わる？
		Encoding:    "json",
		// DisableCaller stops annotating logs with the calling function's file
		// name and line number. By default, all logs are annotated.
		// DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
		// DisableStacktrace completely disables automatic stacktrace capturing. By
		// default, stacktraces are captured for WarnLevel and above logs in
		// development and ErrorLevel and above in production.
		// DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
		// Sampling sets a sampling policy. A nil SamplingConfig disables sampling.
		// Sampling *SamplingConfig `json:"sampling" yaml:"sampling"`
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",   // これを指定しないとMessage自体が表示されない
			LevelKey:       "level", // これを指定しないとLevel自体が表示されない
			TimeKey:        "time",  // これを指定しないとTime自体が表示されない
			NameKey:        "name",
			CallerKey:      "caller",     // これを指定しないとCaller（呼び出し元）自体が表示されない
			FunctionKey:    "func",       // これを指定しないとFunc（呼び出し元の関数）自体が表示されない（Callerに行数出て特定はできるので無くてもよい）
			StacktraceKey:  "stacktrace", // これを指定しないとStacktrace自体が表示されない（Error, Fatalのときに表示される）
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			// EncodeName:       "",
			// ConsoleSeparator: "",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()

	if err != nil {
		panic(err)
	}

	defer logger.Sync() // バッファリングされたログエントリをflushする

	tId := xid.New().String()
	traceId := zap.String("traceId", tId)

	// {"level":"INFO","time":"2021-12-19T16:40:10.293702564+09:00","caller":"zap-sample/main.go:55","func":"main.main","msg":"Hello World","traceId":"c6ve3mhku8rtn7tr1ps0","status":200}
	logger.Info("Hello World", traceId, zap.Int("status", 200))

	// namespace
	// なぜか前のを引き継ぐ。。
	// {"level":"INFO","time":"2021-12-19T16:40:10.293845764+09:00","caller":"zap-sample/main.go:59","func":"main.main","msg":"Hello World","second":{"traceId":"c6ve3mhku8rtn7tr1ps0","status":200}}
	secondLogger := logger.With(zap.Namespace("second"))
	secondLogger.Info("Hello World", traceId, zap.Int("status", 200))

	// object
	req := &request{
		URL:    "/test",
		Listen: addr{"127.0.0.1", 8080},
		Remote: addr{"127.0.0.1", 31200},
	}
	// {"level":"INFO","time":"2021-12-19T16:55:54.877053395+09:00","caller":"zap-sample/main.go:70","func":"main.main","msg":"new request, in nested object","req":{"url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}}
	logger.Info("new request, in nested object", zap.Object("req", req))
	// {"level":"INFO","time":"2021-12-19T16:55:54.877076569+09:00","caller":"zap-sample/main.go:71","func":"main.main","msg":"new request, inline","url":"/test","ip":"127.0.0.1","port":8080,"remote":{"ip":"127.0.0.1","port":31200}}
	logger.Info("new request, inline", zap.Inline(req))

	traceTarget(logger, traceId)

	// sugarLoggerは少し遅いが、InfofやInfowなどの便利なメソッドが利用できる
	sugarLogger := logger.Sugar()

	sugarLogger.Infof("Hello World traceId=%s, status=%d", tId, 200)

	sugarLogger.Infow("Hello World", "traceId", tId, "status", 200)

	func() {
		logger.Info("Hello Goroutine", traceId, zap.Int("status", 200))
	}()

	logger.Debug("Hello Debug", traceId, zap.Int("status", 200)) // これは表示されない
	logger.Warn("Hello Warn", traceId, zap.Int("status", 200))
	logger.Error("Hello Error", traceId, zap.Int("status", 200))
	logger.Fatal("Hello Fatal", traceId, zap.Int("status", 200))
}

func traceTarget(logger *zap.Logger, traceId zap.Field) {
	defer trace(logger, traceId)()
	time.Sleep(1 * time.Second)
}

func trace(logger *zap.Logger, tid zap.Field) func() {
	startTime := time.Now()
	logger.Info("Start Trace", tid)
	return func() {
		logger.Info("End Trace", tid, zap.Int64("elapsed", time.Since(startTime).Nanoseconds()/1000000))
	}
}

type addr struct {
	IP   string
	Port int
}

type request struct {
	URL    string
	Listen addr
	Remote addr
}

func (a addr) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ip", a.IP)
	enc.AddInt("port", a.Port)
	return nil
}

func (r request) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("url", r.URL)
	zap.Inline(r.Listen).AddTo(enc)
	return enc.AddObject("remote", r.Remote)
}
