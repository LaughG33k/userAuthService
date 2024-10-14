package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LaughG33k/userAuthService/iternal/api/auth"
	"github.com/LaughG33k/userAuthService/iternal/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

type App struct {
	serviceProvider *serviceProvider
	cfg             config.AppConfig
	server          *grpc.Server
}

func (a *App) Run() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	tm, canc := context.WithTimeout(context.Background(), 30*time.Second)

	defer canc()

	if err := a.initDeps(tm); err != nil {
		log.Panic(err)
	}

	<-stop

	a.server.GracefulStop()

}

func (a *App) initDeps(ctx context.Context) error {

	deps := []func(context.Context) error{
		a.initConfig,
		a.initGrpc,
		a.initServiceProvider,
		a.initApi,
		a.initServer,
	}

	for _, v := range deps {
		if err := v(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {

	cfg, err := config.Load()

	if err != nil {
		return err
	}

	a.cfg = cfg
	return nil
}

func (a *App) initGrpc(_ context.Context) error {

	recOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			return p.(error)
		}),
	}

	gs := grpc.NewServer(grpc.ChainUnaryInterceptor(recovery.UnaryServerInterceptor(recOpts...)))

	a.server = gs

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {

	a.serviceProvider = newServiceProvider(a.cfg)

	return nil
}

func (a *App) initApi(ctx context.Context) error {

	auth.New(a.serviceProvider.AuthService(ctx), a.server)

	return nil
}

func (a *App) initServer(ctx context.Context) error {

	l, err := net.Listen("tcp", a.cfg.Addr)

	if err != nil {
		return err
	}

	if err := a.server.Serve(l); err != nil {
		return err
	}

	return nil
}
