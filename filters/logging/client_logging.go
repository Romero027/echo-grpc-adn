package logging

import (
	"context"
	"os"
	"time"

	grpc "github.com/Romero027/grpc-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// UnaryClientInterceptor returns a new unary client interceptor that optionally logs the execution of external gRPC calls.
func UnaryClientInterceptor(logger *zap.Logger) grpc.ADNClientProcessor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.ADNInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		logFinalClientLine(logger, startTime, err, "finished client unary call")
		return err
	}
}

func logFinalClientLine(logger *zap.Logger, startTime time.Time, err error, msg string) {
	duration := time.Now().Sub(startTime).Milliseconds()
	logger.Info("finished client unary call", zap.Duration("Duration", time.Microsecond*time.Duration(duration)), zap.Error(err))
}

// func newClientLoggerFields(ctx context.Context, fullMethodString string) []zapcore.Field {
// 	service := path.Dir(fullMethodString)[1:]
// 	method := path.Base(fullMethodString)
// 	return []zapcore.Field{
// 		ClientField,
// 		zap.String("grpc.service", service),
// 		zap.String("grpc.method", method),
// 	}
// }
