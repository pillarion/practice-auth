package main

import (
	"fmt"
	"log/slog"
	"net"

	dgrpc "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	config "github.com/pillarion/practice-auth/internal/adapter/drivers/config/env"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		slog.Warn("failed to get config", "Error", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		slog.Warn("failed to listen", "Error", err)
	}

	slog.Info("config", "config", cfg)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, dgrpc.NewServer())

	slog.Info("server listening at", "address", lis.Addr().String())

	if err = s.Serve(lis); err != nil {
		slog.Warn("failed to serve", "Error", err)
	}
}
