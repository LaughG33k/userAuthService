package auth

import (
	"github.com/LaughG33k/userAuthService/iternal/service"
	"github.com/LaughG33k/userAuthService/pkg/codegen"
	"google.golang.org/grpc"
)

type AuthApi struct {
	codegen.UnimplementedAuthServer
	authService service.Auth
}

func New(authService service.Auth, server *grpc.Server) *AuthApi {
	api := &AuthApi{
		authService: authService,
	}

	codegen.RegisterAuthServer(server, api)

	return api
}
