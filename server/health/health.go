package health

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
)

const serviceName = "proto.API" // taken from compiled protobuf file api.go.pb (line 793)

var server = health.NewServer()

// Register the health service with a gRPC server.
func Register(srv *grpc.Server) { _ = "STUB: not implemented"; return }

// SetServing marks the service as healthy.
func SetServing() { _ = "STUB: not implemented"; return }

// SetNotServing marks the service as unhealthy.
func SetNotServing() { _ = "STUB: not implemented"; return }
