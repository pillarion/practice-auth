package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"

	"github.com/pillarion/practice-auth/internal/adapter/controller/interceptor"
	pbaccess "github.com/pillarion/practice-auth/pkg/access_v1"
	pbauth "github.com/pillarion/practice-auth/pkg/auth_v1"
	pbuser "github.com/pillarion/practice-auth/pkg/user_v1"
	closer "github.com/pillarion/practice-platform/pkg/closer"
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

func (a *App) initGRPCServer(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key)
	if err != nil {
		log.Fatalf("failed to load TLS keys: %v", err)
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
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
		log.Fatalf("failed to load TLS keys: %v", err)
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
		ReadHeaderTimeout: time.Second * 10,
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
		ReadHeaderTimeout: time.Second * 10,
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		slog.Info("Serving swagger file", "path", path)

		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Info("Open swagger file", "path", path)

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		slog.Info("Read swagger file", "path", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Info("Write swagger file", "path", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Info("Finish serving swagger file", "path", path)
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

		slog.Info("GRPC server is stopped")
		return nil
	}()

	go func() error {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			return fmt.Errorf("failed to run HTTP server: %v", err)
		}

		slog.Info("HTTP server is stopped")
		return nil
	}()

	go func() error {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			return fmt.Errorf("failed to run Swagger server: %v", err)
		}

		slog.Info("Swagger server is stopped")
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
	slog.Info("GRPC server is running", "ListenAddress", lAddress)

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	slog.Info("HTTP server is running", "ListenAddress", a.httpServer.Addr)

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
	slog.Info("Swagger server is running", "ListenAddress", a.swaggerServer.Addr)

	err := a.swaggerServer.ListenAndServeTLS(
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Cert,
		a.serviceProvider.config.TLS.Path+a.serviceProvider.Config().TLS.Key,
	)
	if err != nil {
		return err
	}

	return nil
}
