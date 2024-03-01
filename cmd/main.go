package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	dgrpc "github.com/pillarion/practice-auth/internal/adapter/controller/grpc"
	config "github.com/pillarion/practice-auth/internal/adapter/drivers/config/env"
	"github.com/pillarion/practice-auth/internal/adapter/drivers/db/postgresql"
	"github.com/pillarion/practice-auth/internal/core/service/user"
	desc "github.com/pillarion/practice-auth/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	repo, err := postgresql.New(ctx, &cfg.Database)
	if err != nil {
		log.Fatalf("failed to create repo: %v", err)
	}
	us := user.NewService(repo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, dgrpc.NewServer(us))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	slog.Info("server listening at", "address", lis.Addr().String())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
