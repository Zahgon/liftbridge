package server

import (
	"context"

	grpc "google.golang.org/grpc"
)

// addUserContext parses client ID from context and set client ID in context
func addUserContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// AuthzUnaryInterceptor gets user from TLS-authenticated request and add user to ctx
func AuthzUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AuthzStreamInterceptor gets user from TLS-authenticated stream request and add user to ctx
func AuthzStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}
