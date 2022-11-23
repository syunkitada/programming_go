package logger

import (
	"context"

	"github.com/rs/xid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Tracer struct {
	traceId string
}

func NewTracer() (tracer *Tracer) {
	return &Tracer{
		traceId: xid.New().String(),
	}
}

func (self *Tracer) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	enc.AddString("traceId", self.traceId)
	return
}

func (self *Tracer) GetTraceId() (traceId string) {
	return self.traceId
}

func (self *Tracer) Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, append([]zap.Field{zap.Inline(self)}, fields...)...)
}

func (self *Tracer) Info(msg string, fields ...zap.Field) {
	logger.Info(msg, append([]zap.Field{zap.Inline(self)}, fields...)...)
}

func (self *Tracer) Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, append([]zap.Field{zap.Inline(self)}, fields...)...)
}

func (self *Tracer) Error(msg string, fields ...zap.Field) {
	logger.Error(msg, append([]zap.Field{zap.Inline(self)}, fields...)...)
}

func (self *Tracer) Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, append([]zap.Field{zap.Inline(self)}, fields...)...)
}

const tracerContextKey = "tracer"

func ContextWithTracer(parent context.Context, tracer *Tracer) context.Context {
	return context.WithValue(parent, tracerContextKey, tracer)
}

func TracerFromContext(ctx context.Context) *Tracer {
	v := ctx.Value(tracerContextKey)
	tracer, ok := v.(*Tracer)
	if !ok {
		tracer = NewTracer()
		tracer.Warn("tracer was generated, because of tracer is not found in context")
	}
	return tracer
}
