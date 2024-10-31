package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

type loggerKey struct{}

func NewLoggerInterceptor(logger *slog.Logger) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			ctx = context.WithValue(ctx, loggerKey{}, logger.With(
				"procedure", req.Spec().Procedure,
			))

			return next(ctx, req)
		})
	})
}
