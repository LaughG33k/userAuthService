package grpcserver

import (
	"context"
	"net"

	"github.com/LaughG33k/userAuthService/iternal/handler"
	"github.com/LaughG33k/userAuthService/pkg/grpc/codegen/authservice/authservice/codegen"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	server *grpc.Server
	ctx    context.Context
	addr   string
}

func NewServer(ctx context.Context, logger *zap.Logger, addr string, maxConcConns int, authHandler *handler.GrpcAuthHandler) *GrpcServer {

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Error(p.(string))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recoveryOpts...),
		),

		grpc.MaxConcurrentStreams(uint32(maxConcConns)),
		grpc.NumStreamWorkers(uint32(maxConcConns)),
	)

	codegen.RegisterAuthServer(server, authHandler)

	return &GrpcServer{
		ctx:    ctx,
		server: server,
		addr:   addr,
	}

}

func (s *GrpcServer) Start() error {

	l, err := net.Listen("tcp", s.addr)

	if err != nil {
		return err
	}

	if err := s.server.Serve(l); err != nil {
		return err
	}

	return nil
}
