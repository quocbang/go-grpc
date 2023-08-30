package middleware

import (
	"context"
	"time"

	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	requestID := xid.New().String()
	logger := zap.L().With(zap.String("rid", requestID))

	reqLogLevel := []zap.Field{}
	logger.Info("start grpc request..", reqLogLevel...)

	resp, err = handler(ctx, req)
	defer func(start time.Time, responseErr error) {
		respLogField := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.Any("server", info.Server),
			zap.Any("msg", resp),
			zap.Duration("duration", time.Since(start)),
		}
		if responseErr != nil {
			logger.Error("grpc response error", respLogField...)
			return
		}
		logger.Info("grpc server response", respLogField...)
	}(time.Now(), err)

	return resp, err
}
