package main

import (
	"fmt"
	"log/slog"
	"net"

	dgrpc "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		slog.Warn("failed to listen", "Error", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, dgrpc.NewServer())

	slog.Info("server listening at", "address", lis.Addr().String())

	if err = s.Serve(lis); err != nil {
		slog.Warn("failed to serve", "Error", err)
	}
}
