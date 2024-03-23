package app

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"

	"github.com/pillarion/practice-auth/internal/core/tools/logger"
	pbaccess "github.com/pillarion/practice-auth/pkg/access_v1"
	pbauth "github.com/pillarion/practice-auth/pkg/auth_v1"
	pbuser "github.com/pillarion/practice-auth/pkg/user_v1"
	closer "github.com/pillarion/practice-platform/pkg/closer"
)

const (
	readHeaderTimeout = 10 * time.Second
)

// App is the main application struct.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
}

// NewApp initializes a new App.
//
// Takes a context as a parameter.
// Returns a pointer to App and an error.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initSwaggerServer,
		a.initLogger,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init("practice-auth", "info")
}

func (a *App) initGRPCServer(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key)
	if err != nil {
		logger.FatalOnError("failed to load TLS keys", err)
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			a.serviceProvider.interceptor.ValidateInterceptor,
			a.serviceProvider.interceptor.LogInterceptor,
		)),
	)

	reflection.Register(a.grpcServer)

	pbuser.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserServer(ctx))
	pbauth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthServer(ctx))
	pbaccess.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.AccessServer(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	creds, err := credentials.NewServerTLSFromFile(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key)
	if err != nil {
		logger.FatalOnError("failed to load TLS keys", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	grpcAddr := fmt.Sprintf(":%d", a.serviceProvider.Config().GRPC.Port)
	httpAddr := fmt.Sprintf(":%d", a.serviceProvider.Config().HTTP.Port)

	err = pbuser.RegisterUserV1HandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:              httpAddr,
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}
	addr := fmt.Sprintf(":%d", a.serviceProvider.Config().Swagger.Port)

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerFile("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		logger.Debug().Str("path", path).Msg("Serving swagger file")

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Debug().Str("path", path).Msg("Open swagger file")

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		logger.Debug().Str("path", path).Msg("Read swagger file")

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Debug().Str("path", path).Msg("Write swagger file")

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Debug().Str("path", path).Msg("Finish serving swagger file")
	}
}

// Run runs the App.
//
// No parameters.
// Returns an error.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() error {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			return fmt.Errorf("failed to run GRPC server: %v", err)
		}

		logger.Info().Msg("GRPC server is stopped")
		return nil
	}()

	go func() error {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			return fmt.Errorf("failed to run HTTP server: %v", err)
		}

		logger.Info().Msg("HTTP server is stopped")
		return nil
	}()

	go func() error {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			return fmt.Errorf("failed to run Swagger server: %v", err)
		}

		logger.Info().Msg("Swagger server is stopped")
		return nil
	}()

	wg.Wait()

	return nil
}

func (a *App) runGRPCServer() error {
	lAddress := fmt.Sprintf(":%d", a.serviceProvider.Config().GRPC.Port)
	list, err := net.Listen("tcp", lAddress)
	if err != nil {
		return err
	}
	logger.Info().Str("address", lAddress).Msg("GRPC server is running")

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Info().Str("address", a.httpServer.Addr).Msg("HTTP server is running")

	err := a.httpServer.ListenAndServeTLS(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	logger.Info().Str("address", a.swaggerServer.Addr).Msg("Swagger server is running")

	err := a.swaggerServer.ListenAndServeTLS(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key,
	)
	if err != nil {
		return err
	}

	return nil
}
