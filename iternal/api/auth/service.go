package auth

import "google.golang.org/grpc"

type AuthApi struct {
	server *grpc.Server
}

func New(server *grpc.Server) *AuthApi {
	return &AuthApi{
		server: server,
	}
}
