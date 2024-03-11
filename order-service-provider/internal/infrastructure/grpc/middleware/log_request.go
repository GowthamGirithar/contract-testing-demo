package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func logRequest(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)

	return h, err
}
